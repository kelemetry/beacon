/*
Copyright 2017 The Kelemetry Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/kelemetry/beacon/api/v1/controller"
	"github.com/kelemetry/beacon/api/v1/resource"
	"github.com/kelemetry/beacon/api/v1/transport"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

func GetResourceKindByName(kind string) resource.ResourceKind {
	if kind == "deployments" || kind == "deployment" || kind == "deploy" || kind == "de" {
		return resource.PodResourceKind{}
	} else if kind == "pods" || kind == "pod" || kind == "po" {
		return resource.PodResourceKind{}
	} else if kind == "ingresses" || kind == "ing" || kind == "ingress" || kind == "in" {
		return resource.IngressResourceKind{}
	}
	// default service
	return resource.ServiceResourceKind{}
}
func GetTransportByName(name string) transport.Transport {
	if name == "nats" {
		return &transport.NATSTransport{}
	}
	return &transport.StdoutTransport{}
}

func main() {
	var kubeconfig string
	var master string
	var namespace string
	var kind string
	var transportName string
	var configFileName string

	flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	flag.StringVar(&master, "master", "", "master url")
	flag.StringVar(&namespace, "namespace", "develop", "namespace")
	flag.StringVar(&kind, "kind", "pods", "kind")
	flag.StringVar(&transportName, "transport", "", "transport")
	flag.StringVar(&configFileName, "config", "", "Link to transport config file")

	flag.Parse()

	rk := GetResourceKindByName(kind)

	trans := GetTransportByName(transportName)
	defer trans.Close()
	err := trans.Initialize()
	if err != nil {
		panic("could not initialize transport")
	}
	// trans.(transport.NATSTransport).Conn.Publish("kelemetry/beacon", []byte("startup"))
	// trans.(transport.NATSTransport).Conn.Flush()
	// creates the connection
	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		glog.Fatal(err)
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatal(err)
	}

	// create the pod watcher
	//listWatcher := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), rk.GetKind(), namespace, fields.Everything())
	listWatcher := rk.NewListWatchFromClient(clientset, namespace, fields.Everything())

	// create the workqueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// Bind the workqueue to a cache with the help of an informer. This way we make sure that
	// whenever the cache is updated, the pod key is added to the workqueue.
	// Note that when we finally process the item from the workqueue, we might see a newer version
	// of the Pod than the version which was responsible for triggering the update.
	indexer, informer := cache.NewIndexerInformer(listWatcher, rk.MakeRuntimeObject(), 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// IndexerInformer uses a delta queue, therefore for deletes we have to use this
			// key function.
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	controller := controller.NewController(queue, indexer, informer, rk.(resource.ResourceKind), trans.(transport.Transport))
	//controller.SetTransport(trans)

	// We can now warm up the cache for initial synchronization.
	// Let's suppose that we knew about a pod "mypod" on our last run, therefore add it to the cache.
	// If this pod is not there anymore, the controller will be notified about the removal after the
	// cache has synchronized.
	glog.Info("Warm up queue")

	meta := rk.MakeWarmUpRuntimeObject()
	indexer.Add(meta)

	// Now let's start the controller
	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(1, stop)

	// Wait forever
	select {}
}

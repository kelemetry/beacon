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

package resource

import (
	"k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type IngressResourceKind struct {
}

func (rk IngressResourceKind) GetKind() string {
	return "ingresses"
}

func (rk IngressResourceKind) GetName(obj interface{}) string {

	return obj.(*v1beta1.Ingress).GetNamespace() + "/" + obj.(*v1beta1.Ingress).GetName()
}

func (rk IngressResourceKind) GetStatus(obj interface{}) interface{} {
	return obj.(*v1beta1.Ingress).Status
}

func (rk IngressResourceKind) MakeRuntimeObject() runtime.Object {
	return &v1beta1.Ingress{}
}

func (rk IngressResourceKind) MakeWarmUpRuntimeObject() runtime.Object {
	return &v1beta1.Ingress{ObjectMeta: meta_v1.ObjectMeta{
		Name:      "deleteme",
		Namespace: v1.NamespaceDefault,
	}}
}

func (rk IngressResourceKind) NewListWatchFromClient(clientset *kubernetes.Clientset, namespace string, flds fields.Selector) *cache.ListWatch {
	return cache.NewListWatchFromClient(clientset.Extensions().RESTClient(), rk.GetKind(), namespace, flds)
}

package resource

import (
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type PodResourceKind struct {
}

func (rk PodResourceKind) GetKind() string {
	return "pods"
}

func (rk PodResourceKind) GetName(obj interface{}) string {
	return obj.(*v1.Pod).GetName()
}

func (rk PodResourceKind) GetStatus(obj interface{}) interface{} {
	return obj.(*v1.Pod).Status
}

func (rk PodResourceKind) MakeRuntimeObject() runtime.Object {
	return &v1.Pod{}
}

func (rk PodResourceKind) MakeWarmUpRuntimeObject() runtime.Object {
	return &v1.Pod{ObjectMeta: meta_v1.ObjectMeta{
		Name:      "deleteme",
		Namespace: v1.NamespaceDefault,
	}}
}

func (rk PodResourceKind) NewListWatchFromClient(clientset *kubernetes.Clientset, namespace string, flds fields.Selector) *cache.ListWatch {
	return cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), rk.GetKind(), namespace, flds)
}

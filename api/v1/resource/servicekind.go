package resource

import (
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type ServiceResourceKind struct {
}

func (rk ServiceResourceKind) GetKind() string {
	return "services"
}

func (rk ServiceResourceKind) GetName(obj interface{}) string {
	return obj.(*v1.Service).GetName()
}

func (rk ServiceResourceKind) GetStatus(obj interface{}) interface{} {
	return obj.(*v1.Service).Status
}

func (rk ServiceResourceKind) MakeRuntimeObject() runtime.Object {
	return &v1.Service{}
}

func (rk ServiceResourceKind) MakeWarmUpRuntimeObject() runtime.Object {
	return &v1.Service{ObjectMeta: meta_v1.ObjectMeta{
		Name:      "deleteme",
		Namespace: v1.NamespaceDefault,
	}}
}

func (rk ServiceResourceKind) NewListWatchFromClient(clientset *kubernetes.Clientset, namespace string, flds fields.Selector) *cache.ListWatch {
	return cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), rk.GetKind(), namespace, flds)
}

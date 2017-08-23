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
	return obj.(*v1beta1.Ingress).GetName()
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

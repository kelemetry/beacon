package resource

import (
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type ResourceKind interface {
	GetKind() string
	GetName(obj interface{}) string
	GetStatus(obj interface{}) interface{}
	NewListWatchFromClient(clientset *kubernetes.Clientset, namespace string, flds fields.Selector) *cache.ListWatch
	MakeRuntimeObject() runtime.Object
	MakeWarmUpRuntimeObject() runtime.Object
}

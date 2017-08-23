package resource

import (
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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

func (rk PodResourceKind) MakeRuntimeObject(meta meta_v1.ObjectMeta) runtime.Object {
	return &v1.Pod{ObjectMeta: meta}
}

func (rk PodResourceKind) MakeWarmUpRuntimeObject() runtime.Object {
	return &v1.Pod{ObjectMeta: meta_v1.ObjectMeta{
		Name:      "deleteme",
		Namespace: v1.NamespaceDefault,
	}}
}

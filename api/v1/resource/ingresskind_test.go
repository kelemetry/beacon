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
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

func TestIngressResourceKind_GetKind(t *testing.T) {
	tests := []struct {
		name string
		rk   IngressResourceKind
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rk := IngressResourceKind{}
			if got := rk.GetKind(); got != tt.want {
				t.Errorf("IngressResourceKind.GetKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngressResourceKind_GetName(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		rk   IngressResourceKind
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rk := IngressResourceKind{}
			if got := rk.GetName(tt.args.obj); got != tt.want {
				t.Errorf("IngressResourceKind.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngressResourceKind_GetStatus(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		rk   IngressResourceKind
		args args
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rk := IngressResourceKind{}
			if got := rk.GetStatus(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngressResourceKind.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngressResourceKind_MakeRuntimeObject(t *testing.T) {
	tests := []struct {
		name string
		rk   IngressResourceKind
		want runtime.Object
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rk := IngressResourceKind{}
			if got := rk.MakeRuntimeObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngressResourceKind.MakeRuntimeObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngressResourceKind_MakeWarmUpRuntimeObject(t *testing.T) {
	tests := []struct {
		name string
		rk   IngressResourceKind
		want runtime.Object
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rk := IngressResourceKind{}
			if got := rk.MakeWarmUpRuntimeObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngressResourceKind.MakeWarmUpRuntimeObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngressResourceKind_NewListWatchFromClient(t *testing.T) {
	type args struct {
		clientset *kubernetes.Clientset
		namespace string
		flds      fields.Selector
	}
	tests := []struct {
		name string
		rk   IngressResourceKind
		args args
		want *cache.ListWatch
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rk := IngressResourceKind{}
			if got := rk.NewListWatchFromClient(tt.args.clientset, tt.args.namespace, tt.args.flds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IngressResourceKind.NewListWatchFromClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

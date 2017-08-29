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
	"reflect"
	"testing"

	"github.com/kelemetry/beacon/api/v1/resource"
	"github.com/kelemetry/beacon/api/v1/transport"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func TestGetResourceKindByName(t *testing.T) {
	type args struct {
		kind string
	}
	tests := []struct {
		name string
		args args
		want resource.ResourceKind
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetResourceKindByName(tt.args.kind); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResourceKindByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTransportByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want transport.Transport
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTransportByName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransportByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_main(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 	// TODO: Add test cases.
// 	}
// 	for range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			main()
// 		})
// 	}
// }

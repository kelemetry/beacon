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

package controller

import (
	"reflect"
	"testing"

	"github.com/kelemetry/beacon/api/v1/resource"
	"github.com/kelemetry/beacon/api/v1/transport"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

func TestNewController(t *testing.T) {
	type args struct {
		clientset   *kubernetes.Clientset
		namespace   string
		flds        fields.Selector
		ratelimiter workqueue.RateLimiter
		rk          resource.ResourceKind
		transport   transport.Transport
	}
	tests := []struct {
		name string
		args args
		want *Controller
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.clientset, tt.args.namespace, tt.args.flds, tt.args.ratelimiter, tt.args.rk, tt.args.transport); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_IndexerAdd(t *testing.T) {
	type fields struct {
		clientset *kubernetes.Clientset
		indexer   cache.Indexer
		informer  cache.Controller
		queue     workqueue.RateLimitingInterface
		rk        resource.ResourceKind
		t         transport.Transport
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				clientset: tt.fields.clientset,
				indexer:   tt.fields.indexer,
				informer:  tt.fields.informer,
				queue:     tt.fields.queue,
				rk:        tt.fields.rk,
				t:         tt.fields.t,
			}
			c.IndexerAdd(tt.args.obj)
		})
	}
}

func TestController_SetTransport(t *testing.T) {
	type fields struct {
		clientset *kubernetes.Clientset
		indexer   cache.Indexer
		informer  cache.Controller
		queue     workqueue.RateLimitingInterface
		rk        resource.ResourceKind
		t         transport.Transport
	}
	type args struct {
		trans transport.Transport
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				clientset: tt.fields.clientset,
				indexer:   tt.fields.indexer,
				informer:  tt.fields.informer,
				queue:     tt.fields.queue,
				rk:        tt.fields.rk,
				t:         tt.fields.t,
			}
			c.SetTransport(tt.args.trans)
		})
	}
}

func TestController_PrettyJson(t *testing.T) {
	type fields struct {
		clientset *kubernetes.Clientset
		indexer   cache.Indexer
		informer  cache.Controller
		queue     workqueue.RateLimitingInterface
		rk        resource.ResourceKind
		t         transport.Transport
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				clientset: tt.fields.clientset,
				indexer:   tt.fields.indexer,
				informer:  tt.fields.informer,
				queue:     tt.fields.queue,
				rk:        tt.fields.rk,
				t:         tt.fields.t,
			}
			got, err := c.PrettyJson(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.PrettyJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Controller.PrettyJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_processNextItem(t *testing.T) {
	type fields struct {
		clientset *kubernetes.Clientset
		indexer   cache.Indexer
		informer  cache.Controller
		queue     workqueue.RateLimitingInterface
		rk        resource.ResourceKind
		t         transport.Transport
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				clientset: tt.fields.clientset,
				indexer:   tt.fields.indexer,
				informer:  tt.fields.informer,
				queue:     tt.fields.queue,
				rk:        tt.fields.rk,
				t:         tt.fields.t,
			}
			if got := c.processNextItem(); got != tt.want {
				t.Errorf("Controller.processNextItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_SyncToStdout(t *testing.T) {
	type fields struct {
		clientset *kubernetes.Clientset
		indexer   cache.Indexer
		informer  cache.Controller
		queue     workqueue.RateLimitingInterface
		rk        resource.ResourceKind
		t         transport.Transport
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				clientset: tt.fields.clientset,
				indexer:   tt.fields.indexer,
				informer:  tt.fields.informer,
				queue:     tt.fields.queue,
				rk:        tt.fields.rk,
				t:         tt.fields.t,
			}
			if err := c.SyncToStdout(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Controller.SyncToStdout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestController_handleErr(t *testing.T) {
	type fields struct {
		clientset *kubernetes.Clientset
		indexer   cache.Indexer
		informer  cache.Controller
		queue     workqueue.RateLimitingInterface
		rk        resource.ResourceKind
		t         transport.Transport
	}
	type args struct {
		err error
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				clientset: tt.fields.clientset,
				indexer:   tt.fields.indexer,
				informer:  tt.fields.informer,
				queue:     tt.fields.queue,
				rk:        tt.fields.rk,
				t:         tt.fields.t,
			}
			c.handleErr(tt.args.err, tt.args.key)
		})
	}
}

func TestController_Run(t *testing.T) {
	type fields struct {
		clientset *kubernetes.Clientset
		indexer   cache.Indexer
		informer  cache.Controller
		queue     workqueue.RateLimitingInterface
		rk        resource.ResourceKind
		t         transport.Transport
	}
	type args struct {
		threadiness int
		stopCh      chan struct{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				clientset: tt.fields.clientset,
				indexer:   tt.fields.indexer,
				informer:  tt.fields.informer,
				queue:     tt.fields.queue,
				rk:        tt.fields.rk,
				t:         tt.fields.t,
			}
			c.Run(tt.args.threadiness, tt.args.stopCh)
		})
	}
}

func TestController_runWorker(t *testing.T) {
	type fields struct {
		clientset *kubernetes.Clientset
		indexer   cache.Indexer
		informer  cache.Controller
		queue     workqueue.RateLimitingInterface
		rk        resource.ResourceKind
		t         transport.Transport
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				clientset: tt.fields.clientset,
				indexer:   tt.fields.indexer,
				informer:  tt.fields.informer,
				queue:     tt.fields.queue,
				rk:        tt.fields.rk,
				t:         tt.fields.t,
			}
			c.runWorker()
		})
	}
}

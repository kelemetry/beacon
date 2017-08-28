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

package transport

import (
	"github.com/kelemetry/beacon/api/v1/resource"
	"github.com/kelemetry/beacon/api/v1/signal"
)

type Transport interface {
	Initialize() error
	BroadcastDelete(key string, rk resource.ResourceKind)
	BroadcastSync(key string, obj interface{}, rk resource.ResourceKind)
	Close()
	SetSignal(signal.SignalInterface)
}

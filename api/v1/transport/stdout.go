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
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/kelemetry/beacon/api/v1/resource"
	"github.com/kelemetry/beacon/api/v1/signal"
)

type StdoutTransport struct {
	Signal signal.SignalInterface
}

func (s StdoutTransport) Initialize() error {
	return nil
}
func (s StdoutTransport) Close() {

}
func (s StdoutTransport) BroadcastDelete(key string, rk resource.ResourceKind) {
	fmt.Printf("Resource (%s) %s does not exist anymore\n", rk.GetKind(), key)
}

func (s StdoutTransport) BroadcastSync(key string, obj interface{}, rk resource.ResourceKind) {
	stat, _ := s.PrettyJson(rk.GetStatus(obj))
	fmt.Printf("Sync/Add/Update for Resource (%s) %s\n%v\n", rk.GetKind(), rk.GetName(obj), stat)

}

func (s StdoutTransport) PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (s *StdoutTransport) SetSignal(signal signal.SignalInterface) {
	s.Signal = signal
}

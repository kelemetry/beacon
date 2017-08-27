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
	"log"

	"github.com/golang/glog"
	"github.com/kelemetry/beacon/api/v1/resource"
	nats "github.com/nats-io/go-nats"
)

type NATSTransport struct {
	Servers string
	Conn    *nats.Conn
}

func (s *NATSTransport) Initialize() error {
	glog.Info("nats initialize")
	var err error
	s.Conn, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	s.Conn.Publish("kelemetry/beacon", []byte("innitialized"))
	s.Conn.Flush()
	return nil
}
func (s *NATSTransport) Close() {
	if s.Conn != nil {
		s.Conn.Close()
	}
	glog.Info("NATS server closed")
}

func (s *NATSTransport) BroadcastDelete(key string, rk resource.ResourceKind) {
	if s.Conn == nil {
		glog.Error("Connection closed")
	}
	subj := "kelemetry/beacon"
	msg := fmt.Sprintf("NATS Resource (%s) %s does not exist anymore\n", rk.GetKind(), key)

	s.Conn.Publish(subj, []byte(msg))
	s.Conn.Flush()

	if err := s.Conn.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
	fmt.Printf(msg)
}

func (s *NATSTransport) BroadcastSync(obj interface{}, rk resource.ResourceKind) {
	stat, _ := s.PrettyJson(rk.GetStatus(obj))

	msg := fmt.Sprintf("NATS Sync/Add/Update for Resource (%s) %s\n%v\n", rk.GetKind(), rk.GetName(obj), stat)
	s.Conn.Publish("kelemetry/beacon", []byte(msg))
	s.Conn.Flush()

}

func (s *NATSTransport) PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

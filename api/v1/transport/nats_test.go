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
	"testing"

	"github.com/kelemetry/beacon/api/v1/resource"
	"github.com/kelemetry/beacon/api/v1/signal"
	nats "github.com/nats-io/go-nats"
)

func TestNATSTransport_Initialize(t *testing.T) {
	type fields struct {
		Servers string
		Conn    *nats.Conn
		Signal  signal.SignalInterface
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NATSTransport{
				Servers: tt.fields.Servers,
				Conn:    tt.fields.Conn,
				Signal:  tt.fields.Signal,
			}
			if err := s.Initialize(); (err != nil) != tt.wantErr {
				t.Errorf("NATSTransport.Initialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNATSTransport_Close(t *testing.T) {
	type fields struct {
		Servers string
		Conn    *nats.Conn
		Signal  signal.SignalInterface
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NATSTransport{
				Servers: tt.fields.Servers,
				Conn:    tt.fields.Conn,
				Signal:  tt.fields.Signal,
			}
			s.Close()
		})
	}
}

func TestNATSTransport_BroadcastDelete(t *testing.T) {
	type fields struct {
		Servers string
		Conn    *nats.Conn
		Signal  signal.SignalInterface
	}
	type args struct {
		key string
		rk  resource.ResourceKind
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
			s := &NATSTransport{
				Servers: tt.fields.Servers,
				Conn:    tt.fields.Conn,
				Signal:  tt.fields.Signal,
			}
			s.BroadcastDelete(tt.args.key, tt.args.rk)
		})
	}
}

func TestNATSTransport_BroadcastSync(t *testing.T) {
	type fields struct {
		Servers string
		Conn    *nats.Conn
		Signal  signal.SignalInterface
	}
	type args struct {
		key string
		obj interface{}
		rk  resource.ResourceKind
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
			s := &NATSTransport{
				Servers: tt.fields.Servers,
				Conn:    tt.fields.Conn,
				Signal:  tt.fields.Signal,
			}
			s.BroadcastSync(tt.args.key, tt.args.obj, tt.args.rk)
		})
	}
}

func TestNATSTransport_PrettyJson(t *testing.T) {
	type fields struct {
		Servers string
		Conn    *nats.Conn
		Signal  signal.SignalInterface
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
			s := &NATSTransport{
				Servers: tt.fields.Servers,
				Conn:    tt.fields.Conn,
				Signal:  tt.fields.Signal,
			}
			got, err := s.PrettyJson(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NATSTransport.PrettyJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NATSTransport.PrettyJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNATSTransport_SetSignal(t *testing.T) {
	type fields struct {
		Servers string
		Conn    *nats.Conn
		Signal  signal.SignalInterface
	}
	type args struct {
		signal signal.SignalInterface
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
			s := &NATSTransport{
				Servers: tt.fields.Servers,
				Conn:    tt.fields.Conn,
				Signal:  tt.fields.Signal,
			}
			s.SetSignal(tt.args.signal)
		})
	}
}

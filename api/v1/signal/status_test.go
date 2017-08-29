package signal

import (
	"testing"

	"github.com/kelemetry/beacon/api/v1/resource"
)

func TestStatusSignal_SignalDelete(t *testing.T) {
	type args struct {
		key string
		rk  resource.ResourceKind
	}
	tests := []struct {
		name string
		s    *StatusSignal
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatusSignal{}
			if got := s.SignalDelete(tt.args.key, tt.args.rk); got != tt.want {
				t.Errorf("StatusSignal.SignalDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusSignal_SignalSync(t *testing.T) {
	type args struct {
		key string
		obj interface{}
		rk  resource.ResourceKind
	}
	tests := []struct {
		name string
		s    *StatusSignal
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatusSignal{}
			if got := s.SignalSync(tt.args.key, tt.args.obj, tt.args.rk); got != tt.want {
				t.Errorf("StatusSignal.SignalSync() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusSignal_PrettyJson(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		s       *StatusSignal
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StatusSignal{}
			got, err := s.PrettyJson(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("StatusSignal.PrettyJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StatusSignal.PrettyJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

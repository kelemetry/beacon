package transport

import (
	"github.com/kelemetry/beacon/api/v1/resource"
)

type Transport interface {
	Initialize() error
	BroadcastDelete(key string, rk resource.ResourceKind)
	BroadcastSync(obj interface{}, rk resource.ResourceKind)
	Close()
}

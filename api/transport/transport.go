package transport

import (
	"github.com/kelemetry/beacon/api/resource"
)

type Transport interface {
	BroadcastDelete(key string, rk resource.ResourceKind)
	BroadcastSync(obj interface{}, rk resource.ResourceKind)
}

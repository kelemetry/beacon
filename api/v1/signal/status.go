package signal

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/kelemetry/beacon/api/v1/resource"
)

type SignalInterface interface {
	SignalDelete(key string, rk resource.ResourceKind) string
	SignalSync(obj interface{}, rk resource.ResourceKind) string
}

type StatusSignal struct{}

func (s *StatusSignal) SignalDelete(key string, rk resource.ResourceKind) string {
	return fmt.Sprintf("NATS Resource (%s) %s does not exist anymore\n", rk.GetKind(), key)
}

func (s *StatusSignal) SignalSync(obj interface{}, rk resource.ResourceKind) string {
	stat, _ := s.PrettyJson(rk.GetStatus(obj))

	return fmt.Sprintf("NATS Sync/Add/Update for Resource (%s) %s\n%v\n", rk.GetKind(), rk.GetName(obj), stat)
}

func (s *StatusSignal) PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

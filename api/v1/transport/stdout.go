package transport

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/kelemetry/beacon/api/v1/resource"
)

type StdoutTransport struct {
}

func (s StdoutTransport) Initialize() error {
	return nil
}
func (s StdoutTransport) Close() {

}
func (s StdoutTransport) BroadcastDelete(key string, rk resource.ResourceKind) {
	fmt.Printf("Resource (%s) %s does not exist anymore\n", rk.GetKind(), key)
}

func (s StdoutTransport) BroadcastSync(obj interface{}, rk resource.ResourceKind) {
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

package transport

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/kelemetry/beacon/api/resource"
)

type StdoutTransport struct {
}

func (s StdoutTransport) BroadcastDelete(key string, rk resource.ResourceKind) {
	fmt.Printf("Pod %s does not exist anymore\n", key)
}

func (s StdoutTransport) BroadcastSync(obj interface{}, rk resource.ResourceKind) {
	stat, _ := s.PrettyJson(rk.GetStatus(obj))
	fmt.Printf("Sync/Add/Update for Pod %s\n%v\n", rk.GetName(obj), stat)

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

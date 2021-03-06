package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateListenerResponse struct {
	Listener       *ListenerResp `json:"listener,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateListenerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateListenerResponse struct{}"
	}

	return strings.Join([]string{"UpdateListenerResponse", string(data)}, " ")
}

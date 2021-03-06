package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateListenerRequestBody struct {
	Listener *UpdateListenerReq `json:"listener"`
}

func (o UpdateListenerRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateListenerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateListenerRequestBody", string(data)}, " ")
}

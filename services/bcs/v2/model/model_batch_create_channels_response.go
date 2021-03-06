package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreateChannelsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateChannelsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateChannelsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateChannelsResponse", string(data)}, " ")
}

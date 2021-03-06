package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchAddPeersToChannelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddPeersToChannelResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddPeersToChannelResponse struct{}"
	}

	return strings.Join([]string{"BatchAddPeersToChannelResponse", string(data)}, " ")
}

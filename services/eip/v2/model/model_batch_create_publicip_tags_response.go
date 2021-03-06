package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchCreatePublicipTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreatePublicipTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreatePublicipTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicipTagsResponse", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreatePredefineTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePredefineTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePredefineTagsResponse struct{}"
	}

	return strings.Join([]string{"CreatePredefineTagsResponse", string(data)}, " ")
}

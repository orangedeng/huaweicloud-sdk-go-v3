package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListStacksByTagRequest struct {
	Tags *[]string `json:"tags,omitempty"`
}

func (o ListStacksByTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListStacksByTagRequest struct{}"
	}

	return strings.Join([]string{"ListStacksByTagRequest", string(data)}, " ")
}

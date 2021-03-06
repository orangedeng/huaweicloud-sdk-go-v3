package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTagsResponse struct {
	// 标签列表
	Tags           *[]string `json:"tags,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTagsResponse", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListResourcesResponse struct {
	Total          *int32          `json:"total,omitempty"`
	Resources      *[]ResourceInfo `json:"resources,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListResourcesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesResponse", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListInstancesRequest struct {
	Limit       *int64  `json:"limit,omitempty"`
	Offset      *int64  `json:"offset,omitempty"`
	IsTemporary *bool   `json:"is_temporary,omitempty"`
	Label       *string `json:"label,omitempty"`
	Search      *string `json:"search,omitempty"`
	SortDir     *string `json:"sort_dir,omitempty"`
	SortKey     *string `json:"sort_key,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}

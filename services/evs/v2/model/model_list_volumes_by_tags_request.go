package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListVolumesByTagsRequest struct {
	Body *ListVolumesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListVolumesByTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVolumesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListVolumesByTagsRequest", string(data)}, " ")
}

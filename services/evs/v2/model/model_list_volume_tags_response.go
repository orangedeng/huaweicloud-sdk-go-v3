package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListVolumeTagsResponse struct {
	// 所有云硬盘的标签信息
	Tags           map[string][]string `json:"tags,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListVolumeTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVolumeTagsResponse struct{}"
	}

	return strings.Join([]string{"ListVolumeTagsResponse", string(data)}, " ")
}

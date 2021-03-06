package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceCreateImageMetadataRequest struct {
	Body *GlanceCreateImageMetadataRequestBody `json:"body,omitempty"`
}

func (o GlanceCreateImageMetadataRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceCreateImageMetadataRequest struct{}"
	}

	return strings.Join([]string{"GlanceCreateImageMetadataRequest", string(data)}, " ")
}

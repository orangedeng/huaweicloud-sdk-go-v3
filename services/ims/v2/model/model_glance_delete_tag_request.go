package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceDeleteTagRequest struct {
	ImageId string `json:"image_id"`
	Tag     string `json:"tag"`
}

func (o GlanceDeleteTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceDeleteTagRequest struct{}"
	}

	return strings.Join([]string{"GlanceDeleteTagRequest", string(data)}, " ")
}

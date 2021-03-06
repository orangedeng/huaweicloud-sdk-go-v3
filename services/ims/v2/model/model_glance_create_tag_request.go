package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceCreateTagRequest struct {
	ImageId string `json:"image_id"`
	Tag     string `json:"tag"`
}

func (o GlanceCreateTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceCreateTagRequest struct{}"
	}

	return strings.Join([]string{"GlanceCreateTagRequest", string(data)}, " ")
}

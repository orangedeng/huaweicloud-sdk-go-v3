package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExportImageRequest struct {
	ImageId string                  `json:"image_id"`
	Body    *ExportImageRequestBody `json:"body,omitempty"`
}

func (o ExportImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportImageRequest struct{}"
	}

	return strings.Join([]string{"ExportImageRequest", string(data)}, " ")
}

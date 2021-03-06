package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RegisterImageRequest struct {
	ImageId string                    `json:"image_id"`
	Body    *RegisterImageRequestBody `json:"body,omitempty"`
}

func (o RegisterImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RegisterImageRequest struct{}"
	}

	return strings.Join([]string{"RegisterImageRequest", string(data)}, " ")
}

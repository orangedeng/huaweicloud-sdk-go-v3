package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateDataImageRequest struct {
	Body *CreateDataImageRequestBody `json:"body,omitempty"`
}

func (o CreateDataImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDataImageRequest struct{}"
	}

	return strings.Join([]string{"CreateDataImageRequest", string(data)}, " ")
}

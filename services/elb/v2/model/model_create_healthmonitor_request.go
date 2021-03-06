package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateHealthmonitorRequest struct {
	Body *CreateHealthmonitorRequestBody `json:"body,omitempty"`
}

func (o CreateHealthmonitorRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateHealthmonitorRequest struct{}"
	}

	return strings.Join([]string{"CreateHealthmonitorRequest", string(data)}, " ")
}

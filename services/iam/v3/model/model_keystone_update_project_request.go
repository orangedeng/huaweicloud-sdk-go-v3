package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneUpdateProjectRequest struct {
	ProjectId string                            `json:"project_id"`
	Body      *KeystoneUpdateProjectRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateProjectRequest struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateProjectRequest", string(data)}, " ")
}

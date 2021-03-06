package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneValidateTokenRequest struct {
	XSubjectToken string  `json:"X-Subject-Token"`
	Nocatalog     *string `json:"nocatalog,omitempty"`
}

func (o KeystoneValidateTokenRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneValidateTokenRequest struct{}"
	}

	return strings.Join([]string{"KeystoneValidateTokenRequest", string(data)}, " ")
}

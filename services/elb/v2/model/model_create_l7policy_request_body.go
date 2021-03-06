package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateL7policyRequestBody struct {
	L7policy *CreateL7policyReq `json:"l7policy"`
}

func (o CreateL7policyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateL7policyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateL7policyRequestBody", string(data)}, " ")
}

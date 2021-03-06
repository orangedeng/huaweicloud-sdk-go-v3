package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateL7PolicyRequestBody struct {
	L7policy *CreateL7PolicyOption `json:"l7policy"`
}

func (o CreateL7PolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateL7PolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyRequestBody", string(data)}, " ")
}

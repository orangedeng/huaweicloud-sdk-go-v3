package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateL7PolicyRequestBody struct {
	L7policy *UpdateL7PolicyOption `json:"l7policy"`
}

func (o UpdateL7PolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateL7PolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateL7PolicyRequestBody", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateL7RuleRequestBody struct {
	Rule *UpdateL7RuleOption `json:"rule"`
}

func (o UpdateL7RuleRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateL7RuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateL7RuleRequestBody", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteL7RuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteL7RuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteL7RuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteL7RuleResponse", string(data)}, " ")
}

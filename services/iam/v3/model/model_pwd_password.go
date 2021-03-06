package model

import (
	"encoding/json"

	"strings"
)

//
type PwdPassword struct {
	User *PwdPasswordUser `json:"user"`
}

func (o PwdPassword) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PwdPassword struct{}"
	}

	return strings.Join([]string{"PwdPassword", string(data)}, " ")
}

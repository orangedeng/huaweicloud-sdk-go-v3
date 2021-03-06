package model

import (
	"encoding/json"

	"strings"
)

//
type PwdAuth struct {
	Identity *PwdIdentity `json:"identity"`
	Scope    *AuthScope   `json:"scope"`
}

func (o PwdAuth) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PwdAuth struct{}"
	}

	return strings.Join([]string{"PwdAuth", string(data)}, " ")
}

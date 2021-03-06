package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDomainLoginPolicyResponse struct {
	LoginPolicy    *LoginPolicyResult `json:"login_policy,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowDomainLoginPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainLoginPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainLoginPolicyResponse", string(data)}, " ")
}

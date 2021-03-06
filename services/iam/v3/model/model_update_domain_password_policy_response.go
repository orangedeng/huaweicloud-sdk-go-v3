package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDomainPasswordPolicyResponse struct {
	PasswordPolicy *PasswordPolicyResult `json:"password_policy,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UpdateDomainPasswordPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainPasswordPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainPasswordPolicyResponse", string(data)}, " ")
}

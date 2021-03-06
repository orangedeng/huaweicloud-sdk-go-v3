package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateDomainLoginPolicyRequestBody struct {
	LoginPolicy *LoginPolicyOption `json:"login_policy"`
}

func (o UpdateDomainLoginPolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainLoginPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainLoginPolicyRequestBody", string(data)}, " ")
}

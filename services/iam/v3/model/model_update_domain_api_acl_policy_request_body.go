package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateDomainApiAclPolicyRequestBody struct {
	ApiAclPolicy *AclPolicyOption `json:"api_acl_policy"`
}

func (o UpdateDomainApiAclPolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainApiAclPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainApiAclPolicyRequestBody", string(data)}, " ")
}

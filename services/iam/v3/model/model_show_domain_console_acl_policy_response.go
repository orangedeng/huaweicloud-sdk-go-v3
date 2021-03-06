package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDomainConsoleAclPolicyResponse struct {
	ConsoleAclPolicy *AclPolicyResult `json:"console_acl_policy,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o ShowDomainConsoleAclPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainConsoleAclPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainConsoleAclPolicyResponse", string(data)}, " ")
}

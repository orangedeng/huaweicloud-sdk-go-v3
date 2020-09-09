/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainApiAclPolicyRequest struct {
	DomainId string                               `json:"domain_id"`
	Body     *UpdateDomainApiAclPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainApiAclPolicyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateDomainApiAclPolicyRequest", string(data)}, " ")
}
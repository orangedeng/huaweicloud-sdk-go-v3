package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateRequestThrottlingPolicyV2Request struct {
	ProjectId  string              `json:"project_id"`
	InstanceId string              `json:"instance_id"`
	Body       *ThrottleBindingReq `json:"body,omitempty"`
}

func (o AssociateRequestThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"AssociateRequestThrottlingPolicyV2Request", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListRequestThrottlingPoliciesBindedToApiV2Request struct {
	ProjectId    string  `json:"project_id"`
	InstanceId   string  `json:"instance_id"`
	ApiId        string  `json:"api_id"`
	ThrottleId   *string `json:"throttle_id,omitempty"`
	ThrottleName *string `json:"throttle_name,omitempty"`
	EnvId        *string `json:"env_id,omitempty"`
	Offset       *int64  `json:"offset,omitempty"`
	Limit        *int32  `json:"limit,omitempty"`
}

func (o ListRequestThrottlingPoliciesBindedToApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRequestThrottlingPoliciesBindedToApiV2Request struct{}"
	}

	return strings.Join([]string{"ListRequestThrottlingPoliciesBindedToApiV2Request", string(data)}, " ")
}

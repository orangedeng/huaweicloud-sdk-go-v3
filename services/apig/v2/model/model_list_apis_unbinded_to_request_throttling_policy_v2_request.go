package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApisUnbindedToRequestThrottlingPolicyV2Request struct {
	ProjectId  string  `json:"project_id"`
	InstanceId string  `json:"instance_id"`
	ThrottleId string  `json:"throttle_id"`
	EnvId      *string `json:"env_id,omitempty"`
	GroupId    *string `json:"group_id,omitempty"`
	ApiId      *string `json:"api_id,omitempty"`
	ApiName    *string `json:"api_name,omitempty"`
	Offset     *int64  `json:"offset,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
}

func (o ListApisUnbindedToRequestThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisUnbindedToRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"ListApisUnbindedToRequestThrottlingPolicyV2Request", string(data)}, " ")
}

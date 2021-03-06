package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApisV2Request struct {
	ProjectId     string  `json:"project_id"`
	InstanceId    string  `json:"instance_id"`
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	GroupId       *string `json:"group_id,omitempty"`
	ReqProtocol   *string `json:"req_protocol,omitempty"`
	ReqMethod     *string `json:"req_method,omitempty"`
	ReqUri        *string `json:"req_uri,omitempty"`
	AuthType      *string `json:"auth_type,omitempty"`
	EnvId         *string `json:"env_id,omitempty"`
	Type          *int32  `json:"type,omitempty"`
	Offset        *int64  `json:"offset,omitempty"`
	Limit         *int32  `json:"limit,omitempty"`
	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListApisV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisV2Request struct{}"
	}

	return strings.Join([]string{"ListApisV2Request", string(data)}, " ")
}

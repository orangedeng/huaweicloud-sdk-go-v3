package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApisBindedToSignatureKeyV2Request struct {
	ProjectId  string  `json:"project_id"`
	InstanceId string  `json:"instance_id"`
	SignId     string  `json:"sign_id"`
	EnvId      *string `json:"env_id,omitempty"`
	ApiId      *string `json:"api_id,omitempty"`
	ApiName    *string `json:"api_name,omitempty"`
	GroupId    *string `json:"group_id,omitempty"`
	Offset     *int64  `json:"offset,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
}

func (o ListApisBindedToSignatureKeyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisBindedToSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"ListApisBindedToSignatureKeyV2Request", string(data)}, " ")
}

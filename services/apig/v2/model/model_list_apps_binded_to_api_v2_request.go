package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAppsBindedToApiV2Request struct {
	ProjectId  string  `json:"project_id"`
	InstanceId string  `json:"instance_id"`
	ApiId      *string `json:"api_id,omitempty"`
	AppName    *string `json:"app_name,omitempty"`
	AppId      *string `json:"app_id,omitempty"`
	EnvId      *string `json:"env_id,omitempty"`
	Offset     *int64  `json:"offset,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
}

func (o ListAppsBindedToApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAppsBindedToApiV2Request struct{}"
	}

	return strings.Join([]string{"ListAppsBindedToApiV2Request", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateApiGroupV2Request struct {
	ProjectId  string       `json:"project_id"`
	InstanceId string       `json:"instance_id"`
	GroupId    string       `json:"group_id"`
	Body       *ApiGroupReq `json:"body,omitempty"`
}

func (o UpdateApiGroupV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateApiGroupV2Request struct{}"
	}

	return strings.Join([]string{"UpdateApiGroupV2Request", string(data)}, " ")
}

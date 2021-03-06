package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteApiGroupV2Request struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	GroupId    string `json:"group_id"`
}

func (o DeleteApiGroupV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApiGroupV2Request struct{}"
	}

	return strings.Join([]string{"DeleteApiGroupV2Request", string(data)}, " ")
}

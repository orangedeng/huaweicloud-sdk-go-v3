package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAppQuantitiesV2Request struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
}

func (o ListAppQuantitiesV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAppQuantitiesV2Request struct{}"
	}

	return strings.Join([]string{"ListAppQuantitiesV2Request", string(data)}, " ")
}

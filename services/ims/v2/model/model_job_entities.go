package model

import (
	"encoding/json"

	"strings"
)

//
type JobEntities struct {
	// 镜像ID
	ImageId *string `json:"image_id,omitempty"`
}

func (o JobEntities) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}

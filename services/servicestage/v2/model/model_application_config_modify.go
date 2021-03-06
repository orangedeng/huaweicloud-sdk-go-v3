package model

import (
	"encoding/json"

	"strings"
)

type ApplicationConfigModify struct {
	// 环境ID。
	EnvironmentId string                                `json:"environment_id"`
	Configuration *ApplicationConfigModifyConfiguration `json:"configuration"`
}

func (o ApplicationConfigModify) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApplicationConfigModify struct{}"
	}

	return strings.Join([]string{"ApplicationConfigModify", string(data)}, " ")
}

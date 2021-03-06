package model

import (
	"encoding/json"

	"strings"
)

type PolicyUpdate struct {
	// 是否启用策略
	Enabled *bool `json:"enabled,omitempty"`
	// 策略名称
	Name                *string           `json:"name,omitempty"`
	OperationDefinition *PolicyoOdCreate  `json:"operation_definition,omitempty"`
	Trigger             *PolicyTriggerReq `json:"trigger,omitempty"`
}

func (o PolicyUpdate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PolicyUpdate struct{}"
	}

	return strings.Join([]string{"PolicyUpdate", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteScalingConfigRequest struct {
	ScalingConfigurationId string `json:"scaling_configuration_id"`
}

func (o DeleteScalingConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteScalingConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingConfigRequest", string(data)}, " ")
}

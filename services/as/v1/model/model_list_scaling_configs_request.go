package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListScalingConfigsRequest struct {
	ScalingConfigurationName *string `json:"scaling_configuration_name,omitempty"`
	ImageId                  *string `json:"image_id,omitempty"`
	StartNumber              *int32  `json:"start_number,omitempty"`
	Limit                    *int32  `json:"limit,omitempty"`
}

func (o ListScalingConfigsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListScalingConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListScalingConfigsRequest", string(data)}, " ")
}

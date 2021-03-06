package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"

	"strings"
)

// Request Object
type ListScalingActivityLogsRequest struct {
	ScalingGroupId string           `json:"scaling_group_id"`
	StartTime      *sdktime.SdkTime `json:"start_time,omitempty"`
	EndTime        *sdktime.SdkTime `json:"end_time,omitempty"`
	StartNumber    *int32           `json:"start_number,omitempty"`
	Limit          *int32           `json:"limit,omitempty"`
}

func (o ListScalingActivityLogsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListScalingActivityLogsRequest struct{}"
	}

	return strings.Join([]string{"ListScalingActivityLogsRequest", string(data)}, " ")
}

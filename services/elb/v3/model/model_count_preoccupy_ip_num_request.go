package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CountPreoccupyIpNumRequest struct {
	AvailabilityZoneId *[]string `json:"availability_zone_id,omitempty"`
	IpTargetEnable     *bool     `json:"ip_target_enable,omitempty"`
	IpVersion          *int32    `json:"ip_version,omitempty"`
	L7FlavorId         *string   `json:"l7_flavor_id,omitempty"`
	LoadbalancerId     *string   `json:"loadbalancer_id,omitempty"`
}

func (o CountPreoccupyIpNumRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CountPreoccupyIpNumRequest struct{}"
	}

	return strings.Join([]string{"CountPreoccupyIpNumRequest", string(data)}, " ")
}

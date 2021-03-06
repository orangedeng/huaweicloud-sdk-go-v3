package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowLoadBalancerRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadBalancerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerRequest", string(data)}, " ")
}

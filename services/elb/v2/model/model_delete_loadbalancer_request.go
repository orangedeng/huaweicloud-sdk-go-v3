package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteLoadbalancerRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
	Cascade        *bool  `json:"cascade,omitempty"`
}

func (o DeleteLoadbalancerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowLoadbalancersStatusRequest struct {
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadbalancersStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLoadbalancersStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancersStatusRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowHealthmonitorsRequest struct {
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o ShowHealthmonitorsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowHealthmonitorsRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthmonitorsRequest", string(data)}, " ")
}

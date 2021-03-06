package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteHealthmonitorRequest struct {
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o DeleteHealthmonitorRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteHealthmonitorRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthmonitorRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateHealthmonitorResponse struct {
	Healthmonitor  *HealthmonitorResp `json:"healthmonitor,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateHealthmonitorResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateHealthmonitorResponse struct{}"
	}

	return strings.Join([]string{"CreateHealthmonitorResponse", string(data)}, " ")
}

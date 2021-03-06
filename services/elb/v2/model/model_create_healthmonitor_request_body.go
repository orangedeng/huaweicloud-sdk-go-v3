package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateHealthmonitorRequestBody struct {
	Healthmonitor *CreateHealthmonitorReq `json:"healthmonitor"`
}

func (o CreateHealthmonitorRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateHealthmonitorRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHealthmonitorRequestBody", string(data)}, " ")
}

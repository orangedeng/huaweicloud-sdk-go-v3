package model

import (
	"encoding/json"

	"strings"
)

type Job struct {
	Name   *string      `json:"name,omitempty"`
	Path   *string      `json:"path,omitempty"`
	Params *interface{} `json:"params,omitempty"`
}

func (o Job) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}

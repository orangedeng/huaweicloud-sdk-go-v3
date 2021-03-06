package model

import (
	"encoding/json"

	"strings"
)

//
type Config struct {
	SecurityCompliance *SecurityCompliance `json:"security_compliance"`
}

func (o Config) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Config struct{}"
	}

	return strings.Join([]string{"Config", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowSecurityComplianceByOptionResponse struct {
	Config         *ConfigByOption `json:"config,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o KeystoneShowSecurityComplianceByOptionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowSecurityComplianceByOptionResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowSecurityComplianceByOptionResponse", string(data)}, " ")
}

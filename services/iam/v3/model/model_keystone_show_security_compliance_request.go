package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowSecurityComplianceRequest struct {
	DomainId string `json:"domain_id"`
}

func (o KeystoneShowSecurityComplianceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowSecurityComplianceRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowSecurityComplianceRequest", string(data)}, " ")
}

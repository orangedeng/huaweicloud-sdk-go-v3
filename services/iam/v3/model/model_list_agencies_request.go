package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAgenciesRequest struct {
	DomainId      string  `json:"domain_id"`
	TrustDomainId *string `json:"trust_domain_id,omitempty"`
	Name          *string `json:"name,omitempty"`
}

func (o ListAgenciesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAgenciesRequest struct{}"
	}

	return strings.Join([]string{"ListAgenciesRequest", string(data)}, " ")
}

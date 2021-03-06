package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAgencyResponse struct {
	Agency         *AgencyResult `json:"agency,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAgencyResponse struct{}"
	}

	return strings.Join([]string{"CreateAgencyResponse", string(data)}, " ")
}

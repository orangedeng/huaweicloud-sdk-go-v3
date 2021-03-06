package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAgencyRequest struct {
	AgencyId string                   `json:"agency_id"`
	Body     *UpdateAgencyRequestBody `json:"body,omitempty"`
}

func (o UpdateAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAgencyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAgencyRequest", string(data)}, " ")
}

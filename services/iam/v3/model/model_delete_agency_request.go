package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAgencyRequest struct {
	AgencyId string `json:"agency_id"`
}

func (o DeleteAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAgencyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAgencyRequest", string(data)}, " ")
}

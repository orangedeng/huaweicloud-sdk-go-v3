package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAgencyRequest struct {
	AgencyId string `json:"agency_id"`
}

func (o ShowAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAgencyRequest struct{}"
	}

	return strings.Join([]string{"ShowAgencyRequest", string(data)}, " ")
}

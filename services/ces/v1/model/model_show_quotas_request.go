package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowQuotasRequest struct {
	ContentType string `json:"Content-Type"`
}

func (o ShowQuotasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowQuotaRequest struct {
}

func (o ShowQuotaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotaRequest", string(data)}, " ")
}

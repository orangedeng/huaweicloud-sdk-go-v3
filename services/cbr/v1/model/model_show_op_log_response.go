package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowOpLogResponse struct {
	OperationLog   *OperationLog `json:"operation_log,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowOpLogResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowOpLogResponse struct{}"
	}

	return strings.Join([]string{"ShowOpLogResponse", string(data)}, " ")
}

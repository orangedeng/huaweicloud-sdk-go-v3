package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowOpLogRequest struct {
	OperationLogId string `json:"operation_log_id"`
}

func (o ShowOpLogRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowOpLogRequest struct{}"
	}

	return strings.Join([]string{"ShowOpLogRequest", string(data)}, " ")
}

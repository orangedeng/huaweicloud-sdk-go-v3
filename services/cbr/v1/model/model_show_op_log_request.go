/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

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
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowOpLogRequest", string(data)}, " ")
}

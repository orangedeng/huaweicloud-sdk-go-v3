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
type ShowBackupRequest struct {
	BackupId string `json:"backup_id"`
}

func (o ShowBackupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowBackupRequest", string(data)}, " ")
}

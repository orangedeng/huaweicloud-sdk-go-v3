package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowBackupResponse struct {
	Backup         *BackupDetail `json:"backup,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowBackupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBackupResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupResponse", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CopyBackupRequest struct {
	BackupId string              `json:"backup_id"`
	Body     *BackupReplicateReq `json:"body,omitempty"`
}

func (o CopyBackupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CopyBackupRequest struct{}"
	}

	return strings.Join([]string{"CopyBackupRequest", string(data)}, " ")
}

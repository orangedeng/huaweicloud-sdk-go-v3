package model

import (
	"encoding/json"

	"strings"
)

// 恢复请求body
type BackupRestoreReq struct {
	Restore *BackupRestore `json:"restore"`
}

func (o BackupRestoreReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BackupRestoreReq struct{}"
	}

	return strings.Join([]string{"BackupRestoreReq", string(data)}, " ")
}

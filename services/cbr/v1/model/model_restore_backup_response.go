package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RestoreBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreBackupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreBackupResponse struct{}"
	}

	return strings.Join([]string{"RestoreBackupResponse", string(data)}, " ")
}

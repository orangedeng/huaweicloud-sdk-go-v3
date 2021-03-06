package model

import (
	"encoding/json"

	"strings"
)

type CouchDbInfo struct {
	// couchDB用户名称
	User *string `json:"user,omitempty"`
}

func (o CouchDbInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CouchDbInfo struct{}"
	}

	return strings.Join([]string{"CouchDbInfo", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowPermissionRequest struct {
	RoleId string `json:"role_id"`
}

func (o KeystoneShowPermissionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowPermissionRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowPermissionRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowPermissionResponse struct {
	Role           *RoleResult `json:"role,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o KeystoneShowPermissionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowPermissionResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowPermissionResponse", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowUserRequest struct {
	UserId string `json:"user_id"`
}

func (o KeystoneShowUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowUserRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowUserRequest", string(data)}, " ")
}

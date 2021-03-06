package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowUserResponse struct {
	User           *KeystoneShowUserResult `json:"user,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o KeystoneShowUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowUserResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowUserResponse", string(data)}, " ")
}

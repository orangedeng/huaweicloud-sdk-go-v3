package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowGroupRequest struct {
	GroupId string `json:"group_id"`
}

func (o KeystoneShowGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowGroupRequest", string(data)}, " ")
}

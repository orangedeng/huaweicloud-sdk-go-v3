package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteWhitelistRequest struct {
	WhitelistId string `json:"whitelist_id"`
}

func (o DeleteWhitelistRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteWhitelistRequest struct{}"
	}

	return strings.Join([]string{"DeleteWhitelistRequest", string(data)}, " ")
}

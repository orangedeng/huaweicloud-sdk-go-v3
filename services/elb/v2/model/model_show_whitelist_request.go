package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowWhitelistRequest struct {
	WhitelistId string `json:"whitelist_id"`
}

func (o ShowWhitelistRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowWhitelistRequest struct{}"
	}

	return strings.Join([]string{"ShowWhitelistRequest", string(data)}, " ")
}

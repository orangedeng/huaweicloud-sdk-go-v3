package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListRegionsRequest struct {
}

func (o KeystoneListRegionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListRegionsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListRegionsRequest", string(data)}, " ")
}

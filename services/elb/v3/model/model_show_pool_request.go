package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPoolRequest struct {
	PoolId string `json:"pool_id"`
}

func (o ShowPoolRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPoolRequest struct{}"
	}

	return strings.Join([]string{"ShowPoolRequest", string(data)}, " ")
}

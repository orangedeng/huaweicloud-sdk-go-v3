package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPoolResponse struct {
	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	Pool           *Pool   `json:"pool,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPoolResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPoolResponse struct{}"
	}

	return strings.Join([]string{"ShowPoolResponse", string(data)}, " ")
}

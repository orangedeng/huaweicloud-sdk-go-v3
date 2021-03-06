package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowVaultProjectTagResponse struct {
	// 标签列表
	Tags           *[]TagsResp `json:"tags,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowVaultProjectTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVaultProjectTagResponse struct{}"
	}

	return strings.Join([]string{"ShowVaultProjectTagResponse", string(data)}, " ")
}

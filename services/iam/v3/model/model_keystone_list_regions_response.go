package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneListRegionsResponse struct {
	Links *Links `json:"links,omitempty"`
	// 区域信息列表。
	Regions        *[]Region `json:"regions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o KeystoneListRegionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListRegionsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListRegionsResponse", string(data)}, " ")
}

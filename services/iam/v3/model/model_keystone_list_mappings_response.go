package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneListMappingsResponse struct {
	Links *Links `json:"links,omitempty"`
	// 映射信息列表。
	Mappings       *[]MappingResult `json:"mappings,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o KeystoneListMappingsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListMappingsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListMappingsResponse", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListIpGroupsRequest struct {
	Description *[]string `json:"description,omitempty"`
	Id          *[]string `json:"id,omitempty"`
	IpList      *[]string `json:"ip_list,omitempty"`
	Limit       *int32    `json:"limit,omitempty"`
	Marker      *string   `json:"marker,omitempty"`
	Name        *[]string `json:"name,omitempty"`
	PageReverse *bool     `json:"page_reverse,omitempty"`
}

func (o ListIpGroupsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListIpGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListIpGroupsRequest", string(data)}, " ")
}

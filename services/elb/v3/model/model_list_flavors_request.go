package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListFlavorsRequest struct {
	Id          *[]string `json:"id,omitempty"`
	Limit       *int32    `json:"limit,omitempty"`
	Marker      *string   `json:"marker,omitempty"`
	Name        *[]string `json:"name,omitempty"`
	PageReverse *bool     `json:"page_reverse,omitempty"`
	Shared      *bool     `json:"shared,omitempty"`
	Type        *[]string `json:"type,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}

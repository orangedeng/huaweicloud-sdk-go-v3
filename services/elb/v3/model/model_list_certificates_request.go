package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCertificatesRequest struct {
	AdminStateUp *bool     `json:"admin_state_up,omitempty"`
	Description  *[]string `json:"description,omitempty"`
	Domain       *[]string `json:"domain,omitempty"`
	Id           *[]string `json:"id,omitempty"`
	Limit        *int32    `json:"limit,omitempty"`
	Marker       *string   `json:"marker,omitempty"`
	Name         *[]string `json:"name,omitempty"`
	PageReverse  *bool     `json:"page_reverse,omitempty"`
	Type         *[]string `json:"type,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}

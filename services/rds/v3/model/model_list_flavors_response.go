package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListFlavorsResponse struct {
	// 实例规格信息。
	Flavors        *[]Flavor `json:"flavors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}

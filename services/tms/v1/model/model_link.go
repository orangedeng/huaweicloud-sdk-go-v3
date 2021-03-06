package model

import (
	"encoding/json"

	"strings"
)

// API的URL地址。
type Link struct {
	// API的URL地址。
	Href string `json:"href"`
	// self
	Rel string `json:"rel"`
}

func (o Link) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Link struct{}"
	}

	return strings.Join([]string{"Link", string(data)}, " ")
}

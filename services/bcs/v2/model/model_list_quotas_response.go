package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListQuotasResponse struct {
	// 配额信息
	Resources      *[]Resource `json:"resources,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListQuotasResponse", string(data)}, " ")
}

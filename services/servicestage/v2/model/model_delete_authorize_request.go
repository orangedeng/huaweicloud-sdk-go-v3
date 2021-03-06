package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAuthorizeRequest struct {
	Name string `json:"name"`
}

func (o DeleteAuthorizeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAuthorizeRequest struct{}"
	}

	return strings.Join([]string{"DeleteAuthorizeRequest", string(data)}, " ")
}

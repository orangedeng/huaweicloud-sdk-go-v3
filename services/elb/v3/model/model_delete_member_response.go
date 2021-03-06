package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMemberResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMemberResponse struct{}"
	}

	return strings.Join([]string{"DeleteMemberResponse", string(data)}, " ")
}

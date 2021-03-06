package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type GlanceCreateTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o GlanceCreateTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceCreateTagResponse struct{}"
	}

	return strings.Join([]string{"GlanceCreateTagResponse", string(data)}, " ")
}

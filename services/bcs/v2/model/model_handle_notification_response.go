package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type HandleNotificationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o HandleNotificationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HandleNotificationResponse struct{}"
	}

	return strings.Join([]string{"HandleNotificationResponse", string(data)}, " ")
}

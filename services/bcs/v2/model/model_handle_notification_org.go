package model

import (
	"encoding/json"

	"strings"
)

type HandleNotificationOrg struct {
	// 加入的组织
	Name *string `json:"name,omitempty"`
}

func (o HandleNotificationOrg) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HandleNotificationOrg struct{}"
	}

	return strings.Join([]string{"HandleNotificationOrg", string(data)}, " ")
}

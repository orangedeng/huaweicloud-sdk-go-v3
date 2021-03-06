package model

import (
	"encoding/json"

	"strings"
)

// 策略修改body
type PolicyUpdateReq struct {
	Policy *PolicyUpdate `json:"policy"`
}

func (o PolicyUpdateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PolicyUpdateReq struct{}"
	}

	return strings.Join([]string{"PolicyUpdateReq", string(data)}, " ")
}

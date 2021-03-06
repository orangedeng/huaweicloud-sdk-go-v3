package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPoolsResponse struct {
	// 后端云服务器对象组列表
	Pools          *[]PoolResp `json:"pools,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListPoolsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListPoolsResponse", string(data)}, " ")
}

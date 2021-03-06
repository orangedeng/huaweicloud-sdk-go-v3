package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AddMemberResponse struct {
	// 添加备份共享成员响应信息
	Members *[]Member `json:"members,omitempty"`
	//
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o AddMemberResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddMemberResponse struct{}"
	}

	return strings.Join([]string{"AddMemberResponse", string(data)}, " ")
}

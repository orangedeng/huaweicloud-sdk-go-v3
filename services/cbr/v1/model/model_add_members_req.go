package model

import (
	"encoding/json"

	"strings"
)

// 备份共享目标用户的项目id。
type AddMembersReq struct {
	// 列表，待添加备份共享成员的project_id。
	Members []string `json:"members"`
}

func (o AddMembersReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddMembersReq struct{}"
	}

	return strings.Join([]string{"AddMembersReq", string(data)}, " ")
}

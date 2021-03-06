package model

import (
	"encoding/json"

	"strings"
)

// 获取镜像成员列表
type GlanceImageMembers struct {
	// 共享状态。
	Status string `json:"status"`
	// 共享时间，格式为UTC时间。
	CreatedAt string `json:"created_at"`
	// 更新时间，格式为UTC时间。
	UpdatedAt string `json:"updated_at"`
	// 镜像ID。
	ImageId string `json:"image_id"`
	// 成员ID。
	MemberId string `json:"member_id"`
	// 共享视图。
	Schema string `json:"schema"`
}

func (o GlanceImageMembers) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceImageMembers struct{}"
	}

	return strings.Join([]string{"GlanceImageMembers", string(data)}, " ")
}

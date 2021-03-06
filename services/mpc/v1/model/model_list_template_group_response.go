package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTemplateGroupResponse struct {
	// 模板组信息列表。
	TemplateGroupList *[]TemplateGroup `json:"template_group_list,omitempty"`
	// 转码模板组总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTemplateGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateGroupResponse", string(data)}, " ")
}

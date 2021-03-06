package model

import (
	"encoding/json"

	"strings"
)

// 标签
type Tag struct {
	// 功能描述：标签键
	Key *string `json:"key,omitempty"`
	// 功能描述：标签值
	Value *string `json:"value,omitempty"`
}

func (o Tag) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}

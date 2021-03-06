package model

import (
	"encoding/json"

	"strings"
)

type DeleteTagsOption struct {
	// 标签键。
	Key string `json:"key"`
}

func (o DeleteTagsOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTagsOption struct{}"
	}

	return strings.Join([]string{"DeleteTagsOption", string(data)}, " ")
}

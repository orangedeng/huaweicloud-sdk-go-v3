package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTagsRequest struct {
	XRepoAuth string `json:"X-Repo-Auth"`
	Namespace string `json:"namespace"`
	Project   string `json:"project"`
}

func (o ListTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTagsRequest", string(data)}, " ")
}

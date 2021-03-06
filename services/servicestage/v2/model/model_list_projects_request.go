package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectsRequest struct {
	XRepoAuth string `json:"X-Repo-Auth"`
	Namespace string `json:"namespace"`
}

func (o ListProjectsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectsRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBranchesRequest struct {
	XRepoAuth string `json:"X-Repo-Auth"`
	Namespace string `json:"namespace"`
	Project   string `json:"project"`
}

func (o ListBranchesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBranchesRequest struct{}"
	}

	return strings.Join([]string{"ListBranchesRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEnterpriseProjectRequest struct {
	Body *EnterpriseProject `json:"body,omitempty"`
}

func (o CreateEnterpriseProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseProjectRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateCloudServiceCustomPolicyRequest struct {
	Body *CreateCloudServiceCustomPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateCloudServiceCustomPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCloudServiceCustomPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateCloudServiceCustomPolicyRequest", string(data)}, " ")
}

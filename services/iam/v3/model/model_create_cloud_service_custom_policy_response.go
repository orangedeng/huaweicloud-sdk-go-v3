package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateCloudServiceCustomPolicyResponse struct {
	Role           *ServicePolicyRoleResult `json:"role,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CreateCloudServiceCustomPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCloudServiceCustomPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateCloudServiceCustomPolicyResponse", string(data)}, " ")
}

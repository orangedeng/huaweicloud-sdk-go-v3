package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateCloudServiceCustomPolicyResponse struct {
	Role           *ServicePolicyRoleResult `json:"role,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o UpdateCloudServiceCustomPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCloudServiceCustomPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateCloudServiceCustomPolicyResponse", string(data)}, " ")
}

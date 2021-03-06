package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateCertificateRequestBody struct {
	Certificate *UpdateCertificateOption `json:"certificate"`
}

func (o UpdateCertificateRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequestBody", string(data)}, " ")
}

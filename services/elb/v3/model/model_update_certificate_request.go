package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateCertificateRequest struct {
	CertificateId string                        `json:"certificate_id"`
	Body          *UpdateCertificateRequestBody `json:"body,omitempty"`
}

func (o UpdateCertificateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequest struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequest", string(data)}, " ")
}

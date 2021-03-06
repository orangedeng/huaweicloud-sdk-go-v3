package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateCertificateRequest struct {
	Body *CreateCertificateRequestBody `json:"body,omitempty"`
}

func (o CreateCertificateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCertificateRequest struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfDomainNameCertificateV2Request struct {
	ProjectId     string `json:"project_id"`
	InstanceId    string `json:"instance_id"`
	GroupId       string `json:"group_id"`
	CertificateId string `json:"certificate_id"`
}

func (o ShowDetailsOfDomainNameCertificateV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfDomainNameCertificateV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfDomainNameCertificateV2Request", string(data)}, " ")
}

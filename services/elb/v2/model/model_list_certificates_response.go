package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCertificatesResponse struct {
	// SSL证书列表对象
	Certificates *[]CertificateResp `json:"certificates,omitempty"`
	// 证书的个数
	InstanceNum    *int32 `json:"instance_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCertificatesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCertificatesResponse struct{}"
	}

	return strings.Join([]string{"ListCertificatesResponse", string(data)}, " ")
}

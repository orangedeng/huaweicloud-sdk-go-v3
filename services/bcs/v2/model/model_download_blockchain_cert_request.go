package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type DownloadBlockchainCertRequest struct {
	BlockchainId string                                `json:"blockchain_id"`
	OrgName      string                                `json:"org_name"`
	CertType     DownloadBlockchainCertRequestCertType `json:"cert_type"`
}

func (o DownloadBlockchainCertRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DownloadBlockchainCertRequest struct{}"
	}

	return strings.Join([]string{"DownloadBlockchainCertRequest", string(data)}, " ")
}

type DownloadBlockchainCertRequestCertType struct {
	value string
}

type DownloadBlockchainCertRequestCertTypeEnum struct {
	ADMIN DownloadBlockchainCertRequestCertType
	USER  DownloadBlockchainCertRequestCertType
	CA    DownloadBlockchainCertRequestCertType
}

func GetDownloadBlockchainCertRequestCertTypeEnum() DownloadBlockchainCertRequestCertTypeEnum {
	return DownloadBlockchainCertRequestCertTypeEnum{
		ADMIN: DownloadBlockchainCertRequestCertType{
			value: "admin",
		},
		USER: DownloadBlockchainCertRequestCertType{
			value: "user",
		},
		CA: DownloadBlockchainCertRequestCertType{
			value: "ca",
		},
	}
}

func (c DownloadBlockchainCertRequestCertType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DownloadBlockchainCertRequestCertType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

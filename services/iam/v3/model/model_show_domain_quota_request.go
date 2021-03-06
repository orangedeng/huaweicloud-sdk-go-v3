package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ShowDomainQuotaRequest struct {
	DomainId string                      `json:"domain_id"`
	Type     *ShowDomainQuotaRequestType `json:"type,omitempty"`
}

func (o ShowDomainQuotaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainQuotaRequest", string(data)}, " ")
}

type ShowDomainQuotaRequestType struct {
	value string
}

type ShowDomainQuotaRequestTypeEnum struct {
	USER   ShowDomainQuotaRequestType
	GROUP  ShowDomainQuotaRequestType
	IDP    ShowDomainQuotaRequestType
	AGENCY ShowDomainQuotaRequestType
	POLICY ShowDomainQuotaRequestType
}

func GetShowDomainQuotaRequestTypeEnum() ShowDomainQuotaRequestTypeEnum {
	return ShowDomainQuotaRequestTypeEnum{
		USER: ShowDomainQuotaRequestType{
			value: "user",
		},
		GROUP: ShowDomainQuotaRequestType{
			value: "group",
		},
		IDP: ShowDomainQuotaRequestType{
			value: "idp",
		},
		AGENCY: ShowDomainQuotaRequestType{
			value: "agency",
		},
		POLICY: ShowDomainQuotaRequestType{
			value: "policy",
		},
	}
}

func (c ShowDomainQuotaRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowDomainQuotaRequestType) UnmarshalJSON(b []byte) error {
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

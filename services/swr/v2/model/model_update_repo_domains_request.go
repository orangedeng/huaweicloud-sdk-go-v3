package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type UpdateRepoDomainsRequest struct {
	ContentType  UpdateRepoDomainsRequestContentType `json:"Content-Type"`
	Namespace    string                              `json:"namespace"`
	Repository   string                              `json:"repository"`
	AccessDomain string                              `json:"access_domain"`
	Body         *UpdateRepoDomainsRequestBody       `json:"body,omitempty"`
}

func (o UpdateRepoDomainsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRepoDomainsRequest struct{}"
	}

	return strings.Join([]string{"UpdateRepoDomainsRequest", string(data)}, " ")
}

type UpdateRepoDomainsRequestContentType struct {
	value string
}

type UpdateRepoDomainsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 UpdateRepoDomainsRequestContentType
	APPLICATION_JSON             UpdateRepoDomainsRequestContentType
}

func GetUpdateRepoDomainsRequestContentTypeEnum() UpdateRepoDomainsRequestContentTypeEnum {
	return UpdateRepoDomainsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: UpdateRepoDomainsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: UpdateRepoDomainsRequestContentType{
			value: "application/json",
		},
	}
}

func (c UpdateRepoDomainsRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateRepoDomainsRequestContentType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type CinderListQuotasRequest struct {
	TargetProjectId string                       `json:"target_project_id"`
	Usage           CinderListQuotasRequestUsage `json:"usage"`
}

func (o CinderListQuotasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CinderListQuotasRequest struct{}"
	}

	return strings.Join([]string{"CinderListQuotasRequest", string(data)}, " ")
}

type CinderListQuotasRequestUsage struct {
	value string
}

type CinderListQuotasRequestUsageEnum struct {
	TRUE CinderListQuotasRequestUsage
}

func GetCinderListQuotasRequestUsageEnum() CinderListQuotasRequestUsageEnum {
	return CinderListQuotasRequestUsageEnum{
		TRUE: CinderListQuotasRequestUsage{
			value: "true",
		},
	}
}

func (c CinderListQuotasRequestUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CinderListQuotasRequestUsage) UnmarshalJSON(b []byte) error {
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

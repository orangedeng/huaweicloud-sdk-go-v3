package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListEnvironmentsRequest struct {
	Limit   *int32                        `json:"limit,omitempty"`
	Offset  *int32                        `json:"offset,omitempty"`
	OrderBy *string                       `json:"order_by,omitempty"`
	Order   *ListEnvironmentsRequestOrder `json:"order,omitempty"`
}

func (o ListEnvironmentsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEnvironmentsRequest struct{}"
	}

	return strings.Join([]string{"ListEnvironmentsRequest", string(data)}, " ")
}

type ListEnvironmentsRequestOrder struct {
	value string
}

type ListEnvironmentsRequestOrderEnum struct {
	DESC ListEnvironmentsRequestOrder
	ASC  ListEnvironmentsRequestOrder
}

func GetListEnvironmentsRequestOrderEnum() ListEnvironmentsRequestOrderEnum {
	return ListEnvironmentsRequestOrderEnum{
		DESC: ListEnvironmentsRequestOrder{
			value: "desc",
		},
		ASC: ListEnvironmentsRequestOrder{
			value: "asc",
		},
	}
}

func (c ListEnvironmentsRequestOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListEnvironmentsRequestOrder) UnmarshalJSON(b []byte) error {
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

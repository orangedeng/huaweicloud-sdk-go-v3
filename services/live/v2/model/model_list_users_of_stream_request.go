package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListUsersOfStreamRequest struct {
	PlayDomain string                            `json:"play_domain"`
	App        *string                           `json:"app,omitempty"`
	Stream     *string                           `json:"stream,omitempty"`
	Isp        *[]string                         `json:"isp,omitempty"`
	Region     *[]string                         `json:"region,omitempty"`
	Interval   *ListUsersOfStreamRequestInterval `json:"interval,omitempty"`
	StartTime  *string                           `json:"start_time,omitempty"`
	EndTime    *string                           `json:"end_time,omitempty"`
}

func (o ListUsersOfStreamRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListUsersOfStreamRequest struct{}"
	}

	return strings.Join([]string{"ListUsersOfStreamRequest", string(data)}, " ")
}

type ListUsersOfStreamRequestInterval struct {
	value int32
}

type ListUsersOfStreamRequestIntervalEnum struct {
	E_60  ListUsersOfStreamRequestInterval
	E_300 ListUsersOfStreamRequestInterval
}

func GetListUsersOfStreamRequestIntervalEnum() ListUsersOfStreamRequestIntervalEnum {
	return ListUsersOfStreamRequestIntervalEnum{
		E_60: ListUsersOfStreamRequestInterval{
			value: 60,
		}, E_300: ListUsersOfStreamRequestInterval{
			value: 300,
		},
	}
}

func (c ListUsersOfStreamRequestInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListUsersOfStreamRequestInterval) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListEventsRequest struct {
	ContentType string                      `json:"Content-Type"`
	EventType   *ListEventsRequestEventType `json:"event_type,omitempty"`
	EventName   *string                     `json:"event_name,omitempty"`
	From        *int64                      `json:"from,omitempty"`
	To          *int64                      `json:"to,omitempty"`
	Start       *int32                      `json:"start,omitempty"`
	Limit       *int32                      `json:"limit,omitempty"`
}

func (o ListEventsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEventsRequest struct{}"
	}

	return strings.Join([]string{"ListEventsRequest", string(data)}, " ")
}

type ListEventsRequestEventType struct {
	value string
}

type ListEventsRequestEventTypeEnum struct {
	EVENT_SYS    ListEventsRequestEventType
	EVENT_CUSTOM ListEventsRequestEventType
}

func GetListEventsRequestEventTypeEnum() ListEventsRequestEventTypeEnum {
	return ListEventsRequestEventTypeEnum{
		EVENT_SYS: ListEventsRequestEventType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: ListEventsRequestEventType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c ListEventsRequestEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListEventsRequestEventType) UnmarshalJSON(b []byte) error {
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

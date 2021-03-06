package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type KeystoneListEndpointsRequest struct {
	Interface *KeystoneListEndpointsRequestInterface `json:"interface,omitempty"`
	ServiceId *string                                `json:"service_id,omitempty"`
}

func (o KeystoneListEndpointsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListEndpointsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListEndpointsRequest", string(data)}, " ")
}

type KeystoneListEndpointsRequestInterface struct {
	value string
}

type KeystoneListEndpointsRequestInterfaceEnum struct {
	PUBLIC   KeystoneListEndpointsRequestInterface
	INTERNAL KeystoneListEndpointsRequestInterface
	ADMIN    KeystoneListEndpointsRequestInterface
}

func GetKeystoneListEndpointsRequestInterfaceEnum() KeystoneListEndpointsRequestInterfaceEnum {
	return KeystoneListEndpointsRequestInterfaceEnum{
		PUBLIC: KeystoneListEndpointsRequestInterface{
			value: "public",
		},
		INTERNAL: KeystoneListEndpointsRequestInterface{
			value: "internal",
		},
		ADMIN: KeystoneListEndpointsRequestInterface{
			value: "admin",
		},
	}
}

func (c KeystoneListEndpointsRequestInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *KeystoneListEndpointsRequestInterface) UnmarshalJSON(b []byte) error {
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

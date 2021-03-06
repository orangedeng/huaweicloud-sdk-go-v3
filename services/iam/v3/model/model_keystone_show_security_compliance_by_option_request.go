package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type KeystoneShowSecurityComplianceByOptionRequest struct {
	DomainId string                                              `json:"domain_id"`
	Option   KeystoneShowSecurityComplianceByOptionRequestOption `json:"option"`
}

func (o KeystoneShowSecurityComplianceByOptionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowSecurityComplianceByOptionRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowSecurityComplianceByOptionRequest", string(data)}, " ")
}

type KeystoneShowSecurityComplianceByOptionRequestOption struct {
	value string
}

type KeystoneShowSecurityComplianceByOptionRequestOptionEnum struct {
	PASSWORD_REGEX             KeystoneShowSecurityComplianceByOptionRequestOption
	PASSWORD_REGEX_DESCRIPTION KeystoneShowSecurityComplianceByOptionRequestOption
}

func GetKeystoneShowSecurityComplianceByOptionRequestOptionEnum() KeystoneShowSecurityComplianceByOptionRequestOptionEnum {
	return KeystoneShowSecurityComplianceByOptionRequestOptionEnum{
		PASSWORD_REGEX: KeystoneShowSecurityComplianceByOptionRequestOption{
			value: "password_regex",
		},
		PASSWORD_REGEX_DESCRIPTION: KeystoneShowSecurityComplianceByOptionRequestOption{
			value: "password_regex_description",
		},
	}
}

func (c KeystoneShowSecurityComplianceByOptionRequestOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *KeystoneShowSecurityComplianceByOptionRequestOption) UnmarshalJSON(b []byte) error {
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

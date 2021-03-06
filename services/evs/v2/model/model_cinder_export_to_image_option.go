package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// 将云硬盘导出为镜像的请求参数
type CinderExportToImageOption struct {
	// 云硬盘导出镜像的容器类型。  目前支持ami、ari、aki、ovf、bare。默认是bare。
	ContainerFormat *CinderExportToImageOptionContainerFormat `json:"container_format,omitempty"`
	// 云硬盘导出镜像的格式。  目前支持vhd、zvhd、zvhd2、raw、qcow2。默认是vhd。
	DiskFormat *CinderExportToImageOptionDiskFormat `json:"disk_format,omitempty"`
	// 强制导出镜像的标示，默认值是false。  当force标记为false时，云硬盘处于正在使用状态时，不能强制导出镜像。 当force标记为true时，即使云硬盘处于正在使用状态时，仍可以导出镜像。
	Force *bool `json:"force,omitempty"`
	// 云硬盘导出镜像的名称。  名称的长度范围为1～128位。 名称只能包含以下字符：大写字母、小写字母、中文、数字、特殊字符包含“-”、“.”、“_”和空格。
	ImageName string `json:"image_name"`
	// 云硬盘导出镜像的系统类型。目前只支持“windows”和“linux”，默认值是“linux”。  说明： 只有云硬盘的volume_image_metadata信息中无“__os_type”字段且云硬盘状态为“available”时，设置的__os_type才会生效。 如果不传递该参数，则使用默认的“linux”值作为镜像的系统类型。
	OsType *CinderExportToImageOptionOsType `json:"__os_type,omitempty"`
}

func (o CinderExportToImageOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CinderExportToImageOption struct{}"
	}

	return strings.Join([]string{"CinderExportToImageOption", string(data)}, " ")
}

type CinderExportToImageOptionContainerFormat struct {
	value string
}

type CinderExportToImageOptionContainerFormatEnum struct {
	AMI  CinderExportToImageOptionContainerFormat
	ARI  CinderExportToImageOptionContainerFormat
	AKI  CinderExportToImageOptionContainerFormat
	OVF  CinderExportToImageOptionContainerFormat
	BARE CinderExportToImageOptionContainerFormat
}

func GetCinderExportToImageOptionContainerFormatEnum() CinderExportToImageOptionContainerFormatEnum {
	return CinderExportToImageOptionContainerFormatEnum{
		AMI: CinderExportToImageOptionContainerFormat{
			value: "ami",
		},
		ARI: CinderExportToImageOptionContainerFormat{
			value: "ari",
		},
		AKI: CinderExportToImageOptionContainerFormat{
			value: "aki",
		},
		OVF: CinderExportToImageOptionContainerFormat{
			value: "ovf",
		},
		BARE: CinderExportToImageOptionContainerFormat{
			value: "bare",
		},
	}
}

func (c CinderExportToImageOptionContainerFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CinderExportToImageOptionContainerFormat) UnmarshalJSON(b []byte) error {
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

type CinderExportToImageOptionDiskFormat struct {
	value string
}

type CinderExportToImageOptionDiskFormatEnum struct {
	VHD   CinderExportToImageOptionDiskFormat
	ZVHD  CinderExportToImageOptionDiskFormat
	ZVHD2 CinderExportToImageOptionDiskFormat
	RAW   CinderExportToImageOptionDiskFormat
	QCOW2 CinderExportToImageOptionDiskFormat
}

func GetCinderExportToImageOptionDiskFormatEnum() CinderExportToImageOptionDiskFormatEnum {
	return CinderExportToImageOptionDiskFormatEnum{
		VHD: CinderExportToImageOptionDiskFormat{
			value: "vhd",
		},
		ZVHD: CinderExportToImageOptionDiskFormat{
			value: "zvhd",
		},
		ZVHD2: CinderExportToImageOptionDiskFormat{
			value: "zvhd2",
		},
		RAW: CinderExportToImageOptionDiskFormat{
			value: "raw",
		},
		QCOW2: CinderExportToImageOptionDiskFormat{
			value: "qcow2",
		},
	}
}

func (c CinderExportToImageOptionDiskFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CinderExportToImageOptionDiskFormat) UnmarshalJSON(b []byte) error {
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

type CinderExportToImageOptionOsType struct {
	value string
}

type CinderExportToImageOptionOsTypeEnum struct {
	WINDOWS CinderExportToImageOptionOsType
	LINUX   CinderExportToImageOptionOsType
}

func GetCinderExportToImageOptionOsTypeEnum() CinderExportToImageOptionOsTypeEnum {
	return CinderExportToImageOptionOsTypeEnum{
		WINDOWS: CinderExportToImageOptionOsType{
			value: "windows",
		},
		LINUX: CinderExportToImageOptionOsType{
			value: "linux",
		},
	}
}

func (c CinderExportToImageOptionOsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CinderExportToImageOptionOsType) UnmarshalJSON(b []byte) error {
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

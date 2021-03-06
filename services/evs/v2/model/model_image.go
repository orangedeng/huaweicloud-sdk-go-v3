package model

import (
	"encoding/json"

	"strings"
)

type Image struct {
	// 云硬盘导出镜像的容器类型。  目前支持ami、ari、aki、ovf、bare。默认是bare。
	ContainerFormat *string `json:"container_format,omitempty"`
	// 云硬盘导出镜像的格式。  目前支持vhd、zvhd、zvhd2、raw、qcow2。默认是vhd。
	DiskFormat *string `json:"disk_format,omitempty"`
	// 云硬盘描述信息。
	DisplayDescription *string `json:"display_description,omitempty"`
	// 云硬盘ID。
	Id string `json:"id"`
	// 云硬盘导出镜像的ID。
	ImageId string `json:"image_id"`
	// 云硬盘导出镜像的名称
	ImageName string `json:"image_name"`
	// 云硬盘容量。
	Size int32 `json:"size"`
	// 云硬盘导出镜像后的状态，正常值为 “uploading”。
	Status string `json:"status"`
	// 云硬盘更新时间。  时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	UpdatedAt  string      `json:"updated_at"`
	VolumeType *VolumeType `json:"volume_type,omitempty"`
}

func (o Image) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Image struct{}"
	}

	return strings.Join([]string{"Image", string(data)}, " ")
}

/*
    * EVS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

type UpdateSnapshotOption struct {
	// 云硬盘快照描述。最大支持255个字节。
	Description string `json:"description,omitempty"`
	// 云硬盘快照名称。最大支持255个字节。
	Name string `json:"name,omitempty"`
}
/*
    * TMS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// 删除预定义标签请求
type ReqDeletePredefineTag struct {
	// 操作标识（区分大小写）：delete（删除）
	Action string `json:"action"`
	// 标签列表
	Tags []PredefineTagRequest `json:"tags"`
}
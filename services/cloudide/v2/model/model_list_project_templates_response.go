/*
    * cloudide
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Response Object
type ListProjectTemplatesResponse struct {
	// 模板列表
	Templates []ProjectTemplates `json:"templates,omitempty"`
	// 状态
	Status string `json:"status,omitempty"`
}
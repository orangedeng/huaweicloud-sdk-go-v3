/*
    * IAM
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Response Object
type KeystoneListEndpointsResponse struct {
	Links *Links `json:"links,omitempty"`
	// 终端节点信息列表。
	Endpoints []Endpoint `json:"endpoints,omitempty"`
}
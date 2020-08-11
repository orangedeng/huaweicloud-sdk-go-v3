/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// peering对象
type CreateVpcPeeringOption struct {
	// 功能说明：对等连接名称 取值范围：支持1~64个字符
	Name string `json:"name"`
	RequestVpcInfo *VpcInfo `json:"request_vpc_info"`
	AcceptVpcInfo *VpcInfo `json:"accept_vpc_info"`
}
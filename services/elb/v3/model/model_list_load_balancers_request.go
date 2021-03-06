package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListLoadBalancersRequest struct {
	AdminStateUp             *bool     `json:"admin_state_up,omitempty"`
	AvailabilityZoneList     *[]string `json:"availability_zone_list,omitempty"`
	BillingInfo              *[]string `json:"billing_info,omitempty"`
	DeletionProtectionEnable *bool     `json:"deletion_protection_enable,omitempty"`
	Description              *[]string `json:"description,omitempty"`
	Eips                     *[]string `json:"eips,omitempty"`
	EnterpriseProjectId      *[]string `json:"enterprise_project_id,omitempty"`
	Guaranteed               *bool     `json:"guaranteed,omitempty"`
	Id                       *[]string `json:"id,omitempty"`
	IpVersion                *[]int32  `json:"ip_version,omitempty"`
	Ipv6VipAddress           *[]string `json:"ipv6_vip_address,omitempty"`
	Ipv6VipPortId            *[]string `json:"ipv6_vip_port_id,omitempty"`
	Ipv6VipVirsubnetId       *[]string `json:"ipv6_vip_virsubnet_id,omitempty"`
	L4FlavorId               *[]string `json:"l4_flavor_id,omitempty"`
	L4ScaleFlavorId          *[]string `json:"l4_scale_flavor_id,omitempty"`
	L7FlavorId               *[]string `json:"l7_flavor_id,omitempty"`
	L7ScaleFlavorId          *[]string `json:"l7_scale_flavor_id,omitempty"`
	Limit                    *int32    `json:"limit,omitempty"`
	Marker                   *string   `json:"marker,omitempty"`
	MemberAddress            *[]string `json:"member_address,omitempty"`
	MemberDeviceId           *[]string `json:"member_device_id,omitempty"`
	Name                     *[]string `json:"name,omitempty"`
	OperatingStatus          *[]string `json:"operating_status,omitempty"`
	PageReverse              *bool     `json:"page_reverse,omitempty"`
	ProvisioningStatus       *[]string `json:"provisioning_status,omitempty"`
	Publicips                *[]string `json:"publicips,omitempty"`
	VipAddress               *[]string `json:"vip_address,omitempty"`
	VipPortId                *[]string `json:"vip_port_id,omitempty"`
	VipSubnetCidrId          *[]string `json:"vip_subnet_cidr_id,omitempty"`
	VpcId                    *[]string `json:"vpc_id,omitempty"`
}

func (o ListLoadBalancersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLoadBalancersRequest struct{}"
	}

	return strings.Join([]string{"ListLoadBalancersRequest", string(data)}, " ")
}

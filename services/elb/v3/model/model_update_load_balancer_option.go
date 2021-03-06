package model

import (
	"encoding/json"

	"strings"
)

// 更新负载均衡器的消息请求体
type UpdateLoadBalancerOption struct {
	// 负载均衡器名称。
	Name *string `json:"name,omitempty"`
	// 负载均衡器的管理状态。 说明：负载均衡器的管理状态。只支持设定为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 负载均衡器功能说明。
	Description *string `json:"description,omitempty"`
	// 双栈实例对应v6的网络id 。 注： 1.默认为空，只有开启IPv6时才会传入。 2.仅当guaranteed是true的场合，才支持更新。 3.解绑ipv6的情况下，输入null
	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`
	// 负载均衡器所在的子网ID。 注： 1.仅当guaranteed是true的场合，才支持更新。 2.解绑ipv4私网的情况下，输入null
	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`
	// 负载均衡器的虚拟IP。 说明： 1.传入vip_address时必须传入vip_subnet_cidr_id 2.不传入vip_address，自动分配虚拟IP。 3.传入vip_address，需要保证该ip地址在子网中未被占用 注：仅当guaranteed是true的场合，才支持更新。
	VipAddress *string `json:"vip_address,omitempty"`
	// 四层Flavor。 注：1.仅当guaranteed是true的场合，才支持更新。 2.不允许非null变成null，null变成非null，只允许改大，不允许改小。
	L4FlavorId *string `json:"l4_flavor_id,omitempty"`
	// 七层Flavor。 注：1.仅当guaranteed是true的场合，才支持更新。 2.不允许非null变成null，null变成非null；只允许改大，不允许改小。
	L7FlavorId    *string       `json:"l7_flavor_id,omitempty"`
	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`
	// 是否启用跨VPC后端转发，值只允许为true
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`
	// 是否开启删除保护
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`
	// 【描述】下联面网络ID列表，若该字段不指定，在loadbalancer所属的VPC中任意选一个网络id，优选双栈网络。 【约束】 1、所有ID同属一个VPC 2、不允许移除已被ELB使用的子网 3、不支持边缘云子网 4、负载均衡实例不处于ACTIVE状态时，不支持该字段更新
	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`
}

func (o UpdateLoadBalancerOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerOption struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerOption", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

type UpdatePoolOption struct {
	// 后端云服务器组的管理状态；该字段为预留字段，暂未启用。只支持更新为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 后端云服务器组的描述信息
	Description *string `json:"description,omitempty"`
	// 描述：后端云服务器组的负载均衡算法     取值：   1、ROUND_ROBIN：加权轮询算法；   2、LEAST_CONNECTIONS：加权最少连接算法；   3、SOURCE_IP：源IP算法；   4、QUIC_CID：连接ID算法；   约束：   1、当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。   2、只有pool的protocol为QUIC时，才支持QUIC_CID算法。
	LbAlgorithm *string `json:"lb_algorithm,omitempty"`
	// 后端云服务器组的名称。
	Name               *string                             `json:"name,omitempty"`
	SessionPersistence *UpdatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`
	SlowStart          *UpdatePoolSlowStartOption          `json:"slow_start,omitempty"`
	// 是否开启删除保护，默认不开启
	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`
}

func (o UpdatePoolOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePoolOption struct{}"
	}

	return strings.Join([]string{"UpdatePoolOption", string(data)}, " ")
}

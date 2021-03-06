package model

import (
	"encoding/json"

	"strings"
)

// 创建主机组请求
type CreatePoolOption struct {
	// 后端云服务器组的管理状态；该字段为预留字段，暂未启用。只支持更新为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 后端云服务器组的描述信息
	Description *string `json:"description,omitempty"`
	// 描述：后端云服务器组的负载均衡算法     取值：   1、ROUND_ROBIN：加权轮询算法；   2、LEAST_CONNECTIONS：加权最少连接算法；   3、SOURCE_IP：源IP算法；   4、QUIC_CID：连接ID算法；   约束：   1、当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。   2、只有pool的protocol为QUIC时，才支持QUIC_CID算法。
	LbAlgorithm string `json:"lb_algorithm"`
	// 后端云服务器组关联的监听器的ID。listener_id和loadbalancer_id中至少指定一个。
	ListenerId *string `json:"listener_id,omitempty"`
	// 后端云服务器组关联的负载均衡器ID。listener_id和loadbalancer_id中至少指定一个。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	// 后端云服务器组的名称。
	Name *string `json:"name,omitempty"`
	// 后端云服务器组所在的项目ID。
	ProjectId *string `json:"project_id,omitempty"`
	// 描述：后端云服务器组的后端协议。   使用说明：支持TCP、UDP、HTTP和QUIC。 约束： 1、listener的protocol为UDP时，pool的protocol必须为UDP或QUIC；   2、listener的protocol为TCP时pool的protocol必须为TCP；   3、listener的protocol为HTTP或TERMINATED_HTTPS时，pool的protocol必须为HTTP。
	Protocol           string                              `json:"protocol"`
	SessionPersistence *CreatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`
	SlowStart          *CreatePoolSlowStartOption          `json:"slow_start,omitempty"`
	// 是否开启删除保护，默认不开启
	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`
}

func (o CreatePoolOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePoolOption struct{}"
	}

	return strings.Join([]string{"CreatePoolOption", string(data)}, " ")
}

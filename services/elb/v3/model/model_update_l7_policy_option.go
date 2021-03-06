package model

import (
	"encoding/json"

	"strings"
)

type UpdateL7PolicyOption struct {
	// 转发策略的管理状态；该字段为预留字段，暂未启用。默认为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发策略描述信息。
	Description *string `json:"description,omitempty"`
	// 转发策略名称。
	Name *string `json:"name,omitempty"`
	// 转发到的listener的ID，当action为REDIRECT_TO_LISTENER时生效。使用说明：只支持protocol为TERMINATED_HTTPS的listener。不能指定为其他loadbalancer下的listener。当action为REDIRECT_TO_POOL时，不可指定。
	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`
	// 转发到pool的ID。当action为REDIRECT_TO_POOL时生效。使用说明：指定的pool不能是listener的default_pool。不能是其他listener的l7policy使用的pool。当action为REDIRECT_TO_LISTENER时，不可指定。
	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`
}

func (o UpdateL7PolicyOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateL7PolicyOption struct{}"
	}

	return strings.Join([]string{"UpdateL7PolicyOption", string(data)}, " ")
}

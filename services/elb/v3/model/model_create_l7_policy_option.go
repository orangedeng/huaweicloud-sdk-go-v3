package model

import (
	"encoding/json"

	"strings"
)

// 转发策略
type CreateL7PolicyOption struct {
	// 转发策略的转发动作；取值：REDIRECT_TO_POOL：转发到后端云服务器组；REDIRECT_TO_LISTENER：重定向到监听器
	Action string `json:"action"`
	// 转发策略的管理状态；该字段为预留字段，暂未启用。默认为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发策略描述信息。
	Description *string `json:"description,omitempty"`
	// 转发策略对应的监听器ID。当action为REDIRECT_TO_POOL时，只支持创建在PROTOCOL为HTTP或TERMINATED_HTTPS的listener上。 当action为REDIRECT_TO_LISTENER时，只支持创建在PROTOCOL为HTTP的listener上。
	ListenerId string `json:"listener_id"`
	// 转发策略名称。
	Name *string `json:"name,omitempty"`
	// 转发策略的优先级，从1递增，最高100。该字段为预留字段，暂未启用。
	Position *int32 `json:"position,omitempty"`
	// 转发策略所在的项目ID。
	ProjectId *string `json:"project_id,omitempty"`
	// 转发到的listener的ID，当action为REDIRECT_TO_LISTENER时生效。当action为REDIRECT_TO_LISTENER时必选。
	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`
	// 转发到pool的ID。当action为REDIRECT_TO_POOL时生效。使用说明：指定的pool不能是listener的default_pool。不能是其他listener的l7policy使用的pool。当action为REDIRECT_TO_POOL时为必选字段。当action为REDIRECT_TO_LISTENER时，不可指定。
	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`
	// 转发到的url。该字段未启用。
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// 转发策略关联的转发规则对象。详细参考表 l7rule字段说明。rules列表中最多含有2个rule对象，且每个rule的type字段不可相同。
	Rules *[]CreateL7PolicyRuleOption `json:"rules,omitempty"`
}

func (o CreateL7PolicyOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateL7PolicyOption struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyOption", string(data)}, " ")
}

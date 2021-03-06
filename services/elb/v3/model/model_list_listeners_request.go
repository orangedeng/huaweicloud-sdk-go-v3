package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListListenersRequest struct {
	AdminStateUp              *bool     `json:"admin_state_up,omitempty"`
	ClientCaTlsContainerRef   *[]string `json:"client_ca_tls_container_ref,omitempty"`
	ClientTimeout             *[]int32  `json:"client_timeout,omitempty"`
	ConnectionLimit           *[]int32  `json:"connection_limit,omitempty"`
	DefaultPoolId             *[]string `json:"default_pool_id,omitempty"`
	DefaultTlsContainerRef    *[]string `json:"default_tls_container_ref,omitempty"`
	Description               *[]string `json:"description,omitempty"`
	EnableMemberRetry         *bool     `json:"enable_member_retry,omitempty"`
	EnterpriseProjectId       *[]string `json:"enterprise_project_id,omitempty"`
	Http2Enable               *bool     `json:"http2_enable,omitempty"`
	Id                        *[]string `json:"id,omitempty"`
	KeepaliveTimeout          *[]int32  `json:"keepalive_timeout,omitempty"`
	Limit                     *int32    `json:"limit,omitempty"`
	LoadbalancerId            *[]string `json:"loadbalancer_id,omitempty"`
	Marker                    *string   `json:"marker,omitempty"`
	MemberAddress             *[]string `json:"member_address,omitempty"`
	MemberDeviceId            *[]string `json:"member_device_id,omitempty"`
	MemberTimeout             *[]int32  `json:"member_timeout,omitempty"`
	Name                      *[]string `json:"name,omitempty"`
	PageReverse               *bool     `json:"page_reverse,omitempty"`
	Protocol                  *[]string `json:"protocol,omitempty"`
	ProtocolPort              *[]string `json:"protocol_port,omitempty"`
	TlsCiphersPolicy          *[]string `json:"tls_ciphers_policy,omitempty"`
	TransparentClientIpEnable *bool     `json:"transparent_client_ip_enable,omitempty"`
}

func (o ListListenersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateLoadBalancerResponse struct {
	Loadbalancer *LoadBalancer `json:"loadbalancer,omitempty"`
	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLoadBalancerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerResponse", string(data)}, " ")
}

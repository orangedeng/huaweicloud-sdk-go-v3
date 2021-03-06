package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateLoadBalancerResponse struct {
	// 请求ID。  注：自动生成 。
	RequestId      *string       `json:"request_id,omitempty"`
	Loadbalancer   *LoadBalancer `json:"loadbalancer,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateLoadBalancerResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerResponse", string(data)}, " ")
}

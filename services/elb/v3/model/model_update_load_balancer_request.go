package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLoadBalancerRequest Request Object
type UpdateLoadBalancerRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *UpdateLoadBalancerRequestBody `json:"body,omitempty"`
}

func (o UpdateLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerRequest", string(data)}, " ")
}

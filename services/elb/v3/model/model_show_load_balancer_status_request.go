package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoadBalancerStatusRequest Request Object
type ShowLoadBalancerStatusRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadBalancerStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerStatusRequest", string(data)}, " ")
}

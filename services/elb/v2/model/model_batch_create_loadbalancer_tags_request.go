package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateLoadbalancerTagsRequest Request Object
type BatchCreateLoadbalancerTagsRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *BatchCreateLoadbalancerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateLoadbalancerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadbalancerTagsRequest", string(data)}, " ")
}

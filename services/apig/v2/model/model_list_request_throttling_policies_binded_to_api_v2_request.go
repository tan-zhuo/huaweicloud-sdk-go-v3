package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRequestThrottlingPoliciesBindedToApiV2Request Request Object
type ListRequestThrottlingPoliciesBindedToApiV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// API编号
	ApiId string `json:"api_id"`

	// 流控策略的编号
	ThrottleId *string `json:"throttle_id,omitempty"`

	// 流控策略的名称
	ThrottleName *string `json:"throttle_name,omitempty"`

	// 绑定的环境编号
	EnvId *string `json:"env_id,omitempty"`
}

func (o ListRequestThrottlingPoliciesBindedToApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRequestThrottlingPoliciesBindedToApiV2Request struct{}"
	}

	return strings.Join([]string{"ListRequestThrottlingPoliciesBindedToApiV2Request", string(data)}, " ")
}

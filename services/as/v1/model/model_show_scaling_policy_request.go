package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScalingPolicyRequest Request Object
type ShowScalingPolicyRequest struct {

	// 伸缩组ID。
	ScalingPolicyId string `json:"scaling_policy_id"`
}

func (o ShowScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowScalingPolicyRequest", string(data)}, " ")
}

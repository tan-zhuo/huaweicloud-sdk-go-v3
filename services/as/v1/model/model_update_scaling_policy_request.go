package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScalingPolicyRequest Request Object
type UpdateScalingPolicyRequest struct {

	// 伸缩策略ID。
	ScalingPolicyId string `json:"scaling_policy_id"`

	Body *UpdateScalingPolicyOption `json:"body,omitempty"`
}

func (o UpdateScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateScalingPolicyRequest", string(data)}, " ")
}

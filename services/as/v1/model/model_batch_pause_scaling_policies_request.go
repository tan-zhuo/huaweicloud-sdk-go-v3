package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchPauseScalingPoliciesRequest Request Object
type BatchPauseScalingPoliciesRequest struct {
	Body *BatchPauseScalingPoliciesOption `json:"body,omitempty"`
}

func (o BatchPauseScalingPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPauseScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchPauseScalingPoliciesRequest", string(data)}, " ")
}

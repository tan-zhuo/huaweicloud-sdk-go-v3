package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScalingGroupRequest Request Object
type CreateScalingGroupRequest struct {
	Body *CreateScalingGroupOption `json:"body,omitempty"`
}

func (o CreateScalingGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingGroupRequest", string(data)}, " ")
}

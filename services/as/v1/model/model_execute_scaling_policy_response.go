package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteScalingPolicyResponse Response Object
type ExecuteScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPolicyResponse", string(data)}, " ")
}

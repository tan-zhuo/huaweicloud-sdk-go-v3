package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseScalingPolicyResponse Response Object
type PauseScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PauseScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"PauseScalingPolicyResponse", string(data)}, " ")
}

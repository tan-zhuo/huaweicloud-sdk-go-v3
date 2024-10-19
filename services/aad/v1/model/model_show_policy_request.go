package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyRequest Request Object
type ShowPolicyRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`
}

func (o ShowPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyRequest", string(data)}, " ")
}

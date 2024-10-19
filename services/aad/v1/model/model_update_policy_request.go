package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyRequest Request Object
type UpdatePolicyRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *UpdatePolicyRequestBody `json:"body,omitempty"`
}

func (o UpdatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequest", string(data)}, " ")
}

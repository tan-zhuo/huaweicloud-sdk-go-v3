package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyGroupRequest Request Object
type UpdatePolicyGroupRequest struct {

	// 策略组id。
	PolicyGroupId string `json:"policy_group_id"`

	Body *UpdatePolicyGroupReq `json:"body,omitempty"`
}

func (o UpdatePolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyGroupRequest", string(data)}, " ")
}

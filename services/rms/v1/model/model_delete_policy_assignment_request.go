package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyAssignmentRequest Request Object
type DeletePolicyAssignmentRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o DeletePolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyAssignmentRequest", string(data)}, " ")
}

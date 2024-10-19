package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyAssignmentResponse Response Object
type DeletePolicyAssignmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePolicyAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyAssignmentResponse struct{}"
	}

	return strings.Join([]string{"DeletePolicyAssignmentResponse", string(data)}, " ")
}

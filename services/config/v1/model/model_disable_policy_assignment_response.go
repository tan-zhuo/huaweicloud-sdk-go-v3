package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisablePolicyAssignmentResponse Response Object
type DisablePolicyAssignmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisablePolicyAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisablePolicyAssignmentResponse struct{}"
	}

	return strings.Join([]string{"DisablePolicyAssignmentResponse", string(data)}, " ")
}

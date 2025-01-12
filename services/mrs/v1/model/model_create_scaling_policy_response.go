package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScalingPolicyResponse Response Object
type CreateScalingPolicyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingPolicyResponse", string(data)}, " ")
}

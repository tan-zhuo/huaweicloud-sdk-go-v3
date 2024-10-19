package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyRequest Request Object
type CreatePolicyRequest struct {
	Body *CreatePolicyRequestBody `json:"body,omitempty"`
}

func (o CreatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyRequest struct{}"
	}

	return strings.Join([]string{"CreatePolicyRequest", string(data)}, " ")
}

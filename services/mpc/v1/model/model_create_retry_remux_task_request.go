package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRetryRemuxTaskRequest Request Object
type CreateRetryRemuxTaskRequest struct {
	Body *RemuxRetryReq `json:"body,omitempty"`
}

func (o CreateRetryRemuxTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetryRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateRetryRemuxTaskRequest", string(data)}, " ")
}

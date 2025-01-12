package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRemuxTaskRequest Request Object
type CreateRemuxTaskRequest struct {
	Body *CreateRemuxTaskReq `json:"body,omitempty"`
}

func (o CreateRemuxTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateRemuxTaskRequest", string(data)}, " ")
}

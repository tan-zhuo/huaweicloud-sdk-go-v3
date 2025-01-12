package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTakeOverTaskRequest Request Object
type CreateTakeOverTaskRequest struct {
	Body *CreateTakeOverTaskReq `json:"body,omitempty"`
}

func (o CreateTakeOverTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTakeOverTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTakeOverTaskRequest", string(data)}, " ")
}

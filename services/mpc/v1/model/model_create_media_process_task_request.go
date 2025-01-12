package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMediaProcessTaskRequest Request Object
type CreateMediaProcessTaskRequest struct {
	Body *CreateMediaProcessReq `json:"body,omitempty"`
}

func (o CreateMediaProcessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMediaProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateMediaProcessTaskRequest", string(data)}, " ")
}

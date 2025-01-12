package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAnimatedGraphicsTaskRequest Request Object
type CreateAnimatedGraphicsTaskRequest struct {
	Body *CreateAnimatedGraphicsTaskReq `json:"body,omitempty"`
}

func (o CreateAnimatedGraphicsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnimatedGraphicsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAnimatedGraphicsTaskRequest", string(data)}, " ")
}

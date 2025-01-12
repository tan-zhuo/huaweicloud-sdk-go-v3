package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAnimatedGraphicsTaskResponse Response Object
type CreateAnimatedGraphicsTaskResponse struct {

	// 任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAnimatedGraphicsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnimatedGraphicsTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAnimatedGraphicsTaskResponse", string(data)}, " ")
}

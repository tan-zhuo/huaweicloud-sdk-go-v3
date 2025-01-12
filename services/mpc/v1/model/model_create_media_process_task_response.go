package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMediaProcessTaskResponse Response Object
type CreateMediaProcessTaskResponse struct {

	// 任务Id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMediaProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMediaProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateMediaProcessTaskResponse", string(data)}, " ")
}

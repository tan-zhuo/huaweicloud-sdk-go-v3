package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFailedTaskRequest Request Object
type DeleteFailedTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o DeleteFailedTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFailedTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteFailedTaskRequest", string(data)}, " ")
}

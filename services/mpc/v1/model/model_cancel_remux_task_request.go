package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelRemuxTaskRequest Request Object
type CancelRemuxTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o CancelRemuxTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"CancelRemuxTaskRequest", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTaskStatusResult struct {

	// 运行任务id，即报告id
	TaskRunId *int32 `json:"task_run_id,omitempty"`
}

func (o UpdateTaskStatusResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusResult struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusResult", string(data)}, " ")
}

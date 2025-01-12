package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeypairTaskRequest Request Object
type ListKeypairTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ListKeypairTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeypairTaskRequest struct{}"
	}

	return strings.Join([]string{"ListKeypairTaskRequest", string(data)}, " ")
}

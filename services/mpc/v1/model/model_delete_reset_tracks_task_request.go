package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResetTracksTaskRequest Request Object
type DeleteResetTracksTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o DeleteResetTracksTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResetTracksTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteResetTracksTaskRequest", string(data)}, " ")
}

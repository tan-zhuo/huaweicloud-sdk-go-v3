package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopMigrationTaskSyncRequest Request Object
type StopMigrationTaskSyncRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o StopMigrationTaskSyncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskSyncRequest struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskSyncRequest", string(data)}, " ")
}

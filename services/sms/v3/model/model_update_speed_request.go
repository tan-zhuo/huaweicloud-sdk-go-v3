package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSpeedRequest Request Object
type UpdateSpeedRequest struct {

	// 主机迁移任务的ID
	TaskId string `json:"task_id"`

	Body *SpeedLimit `json:"body,omitempty"`
}

func (o UpdateSpeedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSpeedRequest struct{}"
	}

	return strings.Join([]string{"UpdateSpeedRequest", string(data)}, " ")
}

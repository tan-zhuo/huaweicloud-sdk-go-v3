package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConsistencyResultRequest Request Object
type UpdateConsistencyResultRequest struct {

	// 任务id
	TaskId string `json:"task_id"`

	Body *ConsistencyResultRequestBody `json:"body,omitempty"`
}

func (o UpdateConsistencyResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConsistencyResultRequest struct{}"
	}

	return strings.Join([]string{"UpdateConsistencyResultRequest", string(data)}, " ")
}

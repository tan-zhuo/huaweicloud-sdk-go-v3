package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobExecutionRequest Request Object
type DeleteJobExecutionRequest struct {

	// 作业ID。
	JobExecutionId string `json:"job_execution_id"`
}

func (o DeleteJobExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobExecutionRequest struct{}"
	}

	return strings.Join([]string{"DeleteJobExecutionRequest", string(data)}, " ")
}

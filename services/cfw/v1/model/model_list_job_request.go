package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobRequest Request Object
type ListJobRequest struct {

	// 任务ID
	JobId string `json:"job_id"`
}

func (o ListJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobRequest struct{}"
	}

	return strings.Join([]string{"ListJobRequest", string(data)}, " ")
}

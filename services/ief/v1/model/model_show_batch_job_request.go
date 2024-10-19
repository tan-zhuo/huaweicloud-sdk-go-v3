package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchJobRequest Request Object
type ShowBatchJobRequest struct {

	// 批量处理作业ID
	JobId string `json:"job_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o ShowBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchJobRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchJobRequest", string(data)}, " ")
}

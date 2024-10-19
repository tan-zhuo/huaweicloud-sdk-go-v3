package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetSubJobDetail 子任务详情
type GetSubJobDetail struct {

	// 子job的状态。success：成功。running：运行中。failed：失败。waiting：等待执行。
	Status *string `json:"status,omitempty"`

	// job的ID。
	JobId *string `json:"job_id,omitempty"`

	// 子job的类型。
	JobType *string `json:"job_type,omitempty"`

	// job开始时间。UTC时间，格式：'2016-01-02 15:04:05'
	BeginTime *string `json:"begin_time,omitempty"`

	// job结束时间。UTC时间，格式：'2016-01-02 15:04:05'
	EndTime *string `json:"end_time,omitempty"`

	// job执行失败时的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// job执行失败时的错误原因
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o GetSubJobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetSubJobDetail struct{}"
	}

	return strings.Join([]string{"GetSubJobDetail", string(data)}, " ")
}

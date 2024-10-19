package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionStatisticsResponseBody 函数指标
type ListFunctionStatisticsResponseBody struct {

	// 调用次数
	Count *[]SlaReportsValue `json:"count,omitempty"`

	// 平均时延，单位毫秒
	Duration *[]SlaReportsValue `json:"duration,omitempty"`

	// 错误次数
	FailCount *[]SlaReportsValue `json:"fail_count,omitempty"`

	// 最大时延，单位毫秒
	MaxDuration *[]SlaReportsValue `json:"max_duration,omitempty"`

	// 最小时延，单位毫秒
	MinDuration *[]SlaReportsValue `json:"min_duration,omitempty"`

	// 被拒绝次数
	RejectCount *[]SlaReportsValue `json:"reject_count,omitempty"`

	// 函数错误次数
	FunctionErrorCount *[]SlaReportsValue `json:"function_error_count,omitempty"`

	// 系统错误次数
	SystemErrorCount *[]SlaReportsValue `json:"system_error_count,omitempty"`

	// 预留实例指标
	ReservedInstanceNum *[]SlaReportsValue `json:"reserved_instance_num,omitempty"`

	// 弹性实例指标
	ConcurrencyNum *[]SlaReportsValue `json:"concurrency_num,omitempty"`
}

func (o ListFunctionStatisticsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionStatisticsResponseBody struct{}"
	}

	return strings.Join([]string{"ListFunctionStatisticsResponseBody", string(data)}, " ")
}

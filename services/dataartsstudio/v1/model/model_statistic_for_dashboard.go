package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatisticForDashboard struct {

	// 统计时间点(13位时间戳)
	LogTime *int32 `json:"log_time,omitempty"`

	// 调用总量
	CallNum *int32 `json:"call_num,omitempty"`

	// 成功调用量(取数成功)
	SuccessNum *int32 `json:"success_num,omitempty"`

	// 失败调用量(取数失败)
	FailNum *int32 `json:"fail_num,omitempty"`

	// 合法调用量(通过校验)
	LegalNum *int32 `json:"legal_num,omitempty"`

	// 非法调用量(无法通过校验)
	IllegalNum *int32 `json:"illegal_num,omitempty"`

	// 请求平均时长
	CostTimeAvg float32 `json:"cost_time_avg,omitempty"`

	// 成功请求平均时长
	SuccessCostTimeAvg float32 `json:"success_cost_time_avg,omitempty"`

	// 失败请求平均时长
	FailCostTimeAvg float32 `json:"fail_cost_time_avg,omitempty"`

	// 成功率
	SuccessRate float32 `json:"success_rate,omitempty"`

	// 失败率
	FailRate float32 `json:"fail_rate,omitempty"`

	// 合法率
	LegalRate float32 `json:"legal_rate,omitempty"`

	// 非法率
	IllegalRate float32 `json:"illegal_rate,omitempty"`
}

func (o StatisticForDashboard) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticForDashboard struct{}"
	}

	return strings.Join([]string{"StatisticForDashboard", string(data)}, " ")
}

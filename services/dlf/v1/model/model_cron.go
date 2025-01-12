package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Cron struct {
	StartTime *string `json:"startTime,omitempty"`

	EndTime *string `json:"endTime,omitempty"`

	// Cron表达式
	Expression *string `json:"expression,omitempty"`

	// 是否依赖本作业上一个运行周期任务的执行结果
	DependPrePeriod *bool `json:"dependPrePeriod,omitempty"`

	DependJobs *DependJob `json:"dependJobs,omitempty"`
}

func (o Cron) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cron struct{}"
	}

	return strings.Join([]string{"Cron", string(data)}, " ")
}

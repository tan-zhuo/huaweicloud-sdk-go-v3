package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWeeklyReportsRequest Request Object
type ListWeeklyReportsRequest struct {

	// 每周的起始时间
	PeriodStartDate *string `json:"period_start_date,omitempty"`
}

func (o ListWeeklyReportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWeeklyReportsRequest struct{}"
	}

	return strings.Join([]string{"ListWeeklyReportsRequest", string(data)}, " ")
}

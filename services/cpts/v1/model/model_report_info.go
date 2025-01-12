package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportInfo struct {
	Brokens *ReportbrokensInfo `json:"brokens,omitempty"`

	Details *ReportdetailsInfo `json:"details,omitempty"`

	Outline *ReportoutlineInfo `json:"outline,omitempty"`

	// 响应时间分布
	Rtproportion *string `json:"rtproportion,omitempty"`

	TaskInfo *ReportTaskInfo `json:"taskInfo,omitempty"`

	// 响应时间分布
	RespTimeRange *interface{} `json:"respTimeRange,omitempty"`
}

func (o ReportInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportInfo struct{}"
	}

	return strings.Join([]string{"ReportInfo", string(data)}, " ")
}

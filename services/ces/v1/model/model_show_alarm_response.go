package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmResponse Response Object
type ShowAlarmResponse struct {

	// 告警对象列表。
	MetricAlarms   *[]MetricAlarms `json:"metric_alarms,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmResponse", string(data)}, " ")
}

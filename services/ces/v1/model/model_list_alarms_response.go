package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmsResponse Response Object
type ListAlarmsResponse struct {

	// 告警对象列表。
	MetricAlarms *[]MetricAlarms `json:"metric_alarms,omitempty"`

	MetaData       *MetaData `json:"meta_data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmsResponse", string(data)}, " ")
}

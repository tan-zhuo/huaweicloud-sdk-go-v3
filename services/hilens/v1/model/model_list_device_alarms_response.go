package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeviceAlarmsResponse Response Object
type ListDeviceAlarmsResponse struct {

	// 满足条件的设备告警总数
	Total *int32 `json:"total,omitempty"`

	Data           *[]GetDeviceAlarmArrayObject `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListDeviceAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeviceAlarmsResponse struct{}"
	}

	return strings.Join([]string{"ListDeviceAlarmsResponse", string(data)}, " ")
}

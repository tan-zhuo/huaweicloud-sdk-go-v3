package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHealthMonitorRequest Request Object
type DeleteHealthMonitorRequest struct {

	// 健康检查ID。
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o DeleteHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthMonitorRequest", string(data)}, " ")
}

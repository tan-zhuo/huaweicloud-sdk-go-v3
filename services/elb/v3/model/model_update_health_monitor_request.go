package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHealthMonitorRequest Request Object
type UpdateHealthMonitorRequest struct {

	// 健康检查ID
	HealthmonitorId string `json:"healthmonitor_id"`

	Body *UpdateHealthMonitorRequestBody `json:"body,omitempty"`
}

func (o UpdateHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorRequest", string(data)}, " ")
}

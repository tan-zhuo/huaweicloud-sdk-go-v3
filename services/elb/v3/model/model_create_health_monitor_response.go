package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHealthMonitorResponse Response Object
type CreateHealthMonitorResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	Healthmonitor  *HealthMonitor `json:"healthmonitor,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateHealthMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthMonitorResponse struct{}"
	}

	return strings.Join([]string{"CreateHealthMonitorResponse", string(data)}, " ")
}

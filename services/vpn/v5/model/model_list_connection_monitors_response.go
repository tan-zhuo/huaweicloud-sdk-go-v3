package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionMonitorsResponse Response Object
type ListConnectionMonitorsResponse struct {
	ConnectionMonitors *[]ConnectionMonitorInfo `json:"connection_monitors,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListConnectionMonitorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionMonitorsResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionMonitorsResponse", string(data)}, " ")
}

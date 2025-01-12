package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListActiveOrHistoryAlarmsResponse Response Object
type ListActiveOrHistoryAlarmsResponse struct {

	// 告警信息
	Events *[]Events `json:"events,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListActiveOrHistoryAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActiveOrHistoryAlarmsResponse struct{}"
	}

	return strings.Join([]string{"ListActiveOrHistoryAlarmsResponse", string(data)}, " ")
}

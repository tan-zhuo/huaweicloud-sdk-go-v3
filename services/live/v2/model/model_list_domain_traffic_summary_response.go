package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainTrafficSummaryResponse Response Object
type ListDomainTrafficSummaryResponse struct {

	// 域名对应的流量汇总列表。
	TrafficList *[]TrafficSummaryData `json:"traffic_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDomainTrafficSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainTrafficSummaryResponse struct{}"
	}

	return strings.Join([]string{"ListDomainTrafficSummaryResponse", string(data)}, " ")
}

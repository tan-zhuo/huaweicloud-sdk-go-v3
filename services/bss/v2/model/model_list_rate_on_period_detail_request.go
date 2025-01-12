package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRateOnPeriodDetailRequest Request Object
type ListRateOnPeriodDetailRequest struct {
	Body *RateOnPeriodReq `json:"body,omitempty"`
}

func (o ListRateOnPeriodDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRateOnPeriodDetailRequest struct{}"
	}

	return strings.Join([]string{"ListRateOnPeriodDetailRequest", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMetricNotifyRuleResponse Response Object
type DeleteMetricNotifyRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMetricNotifyRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMetricNotifyRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteMetricNotifyRuleResponse", string(data)}, " ")
}

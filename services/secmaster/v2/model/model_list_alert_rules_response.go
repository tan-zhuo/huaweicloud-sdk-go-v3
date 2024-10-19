package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRulesResponse Response Object
type ListAlertRulesResponse struct {

	// 总数量。Total count.
	Count *int64 `json:"count,omitempty"`

	// 告警模型。Alert rules.
	Records *[]AlertRule `json:"records,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAlertRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAlertRulesResponse", string(data)}, " ")
}

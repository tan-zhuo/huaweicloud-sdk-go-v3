package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlarmRulesRequest Request Object
type CreateAlarmRulesRequest struct {
	Body *PostAlarmsReqV2 `json:"body,omitempty"`
}

func (o CreateAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"CreateAlarmRulesRequest", string(data)}, " ")
}

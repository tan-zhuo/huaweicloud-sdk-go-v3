package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAlarmRuleRequest Request Object
type AddAlarmRuleRequest struct {
	Body *AlarmRuleParam `json:"body,omitempty"`
}

func (o AddAlarmRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAlarmRuleRequest struct{}"
	}

	return strings.Join([]string{"AddAlarmRuleRequest", string(data)}, " ")
}

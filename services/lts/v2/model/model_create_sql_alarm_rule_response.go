package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlAlarmRuleResponse Response Object
type CreateSqlAlarmRuleResponse struct {

	// 告警规则id
	SqlAlarmRuleId *string `json:"sql_alarm_rule_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSqlAlarmRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlAlarmRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlAlarmRuleResponse", string(data)}, " ")
}

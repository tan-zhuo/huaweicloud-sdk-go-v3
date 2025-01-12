package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlAlarmRulesResponse Response Object
type ListSqlAlarmRulesResponse struct {

	// SQL告警
	SqlAlarmRules  *[]SqlAlarmRuleRespList `json:"sql_alarm_rules,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListSqlAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSqlAlarmRulesResponse", string(data)}, " ")
}

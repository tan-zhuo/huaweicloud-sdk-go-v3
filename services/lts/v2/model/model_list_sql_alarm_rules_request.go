package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlAlarmRulesRequest Request Object
type ListSqlAlarmRulesRequest struct {
}

func (o ListSqlAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlAlarmRulesRequest", string(data)}, " ")
}

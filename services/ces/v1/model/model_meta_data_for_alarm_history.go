package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetaDataForAlarmHistory 查询告警历史返回的总条数。
type MetaDataForAlarmHistory struct {

	// 查询告警历史返回的总条数。
	Total int32 `json:"total"`
}

func (o MetaDataForAlarmHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaDataForAlarmHistory struct{}"
	}

	return strings.Join([]string{"MetaDataForAlarmHistory", string(data)}, " ")
}

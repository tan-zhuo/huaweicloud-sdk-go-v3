package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreTablesNewRequestBody struct {

	// 恢复时间戳
	RestoreTime int64 `json:"restore_time"`

	// 表信息
	RestoreTables []RestoreDatabasesInfoNew `json:"restore_tables"`

	// 是否使用极速恢复，可先根据“获取实例是否能使用极速恢复”接口判断本次恢复是否能使用极速恢复。 如果实例使用了XA事务，采用极速恢复的方式会导致恢复失败！
	IsFastRestore *bool `json:"is_fast_restore,omitempty"`
}

func (o RestoreTablesNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTablesNewRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreTablesNewRequestBody", string(data)}, " ")
}

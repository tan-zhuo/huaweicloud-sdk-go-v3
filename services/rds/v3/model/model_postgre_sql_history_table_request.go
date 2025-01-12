package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PostgreSqlHistoryTableRequest 查询可恢复表的请求信息
type PostgreSqlHistoryTableRequest struct {

	// 实例ID集合
	InstanceIds []string `json:"instance_ids"`

	// 恢复时间点
	RestoreTime int64 `json:"restore_time"`

	// 数据库名，模糊查询
	DatabaseNameLike *string `json:"database_name_like,omitempty"`

	// 数据库表名，模糊查询
	TableNameLike *string `json:"table_name_like,omitempty"`

	// 实例名称，模糊查询
	InstanceNameLike *string `json:"instance_name_like,omitempty"`
}

func (o PostgreSqlHistoryTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgreSqlHistoryTableRequest struct{}"
	}

	return strings.Join([]string{"PostgreSqlHistoryTableRequest", string(data)}, " ")
}

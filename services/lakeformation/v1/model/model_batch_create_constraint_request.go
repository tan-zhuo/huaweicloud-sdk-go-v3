package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateConstraintRequest Request Object
type BatchCreateConstraintRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	Body *TableConstraintsInput `json:"body,omitempty"`
}

func (o BatchCreateConstraintRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateConstraintRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateConstraintRequest", string(data)}, " ")
}

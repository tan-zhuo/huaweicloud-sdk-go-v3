package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPartitionsRequest Request Object
type AddPartitionsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	Body *AddPartitionInput `json:"body,omitempty"`
}

func (o AddPartitionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPartitionsRequest struct{}"
	}

	return strings.Join([]string{"AddPartitionsRequest", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartitionNamesRequest Request Object
type ListPartitionNamesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	// 过滤条件表达式
	Filter *string `json:"filter,omitempty"`

	// 查询返回条数
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID
	Marker *string `json:"marker,omitempty"`

	// 是否查询上一页
	ReversePage *bool `json:"reverse_page,omitempty"`
}

func (o ListPartitionNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartitionNamesRequest struct{}"
	}

	return strings.Join([]string{"ListPartitionNamesRequest", string(data)}, " ")
}

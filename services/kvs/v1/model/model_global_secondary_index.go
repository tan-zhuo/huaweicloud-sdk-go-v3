package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GlobalSecondaryIndex struct {

	// 二级索引名称，表内唯一。
	IndexName string `bson:"index_name"`

	// 分区键字段名数组，顺序组合。
	ShardKeyFields []Field `bson:"shard_key_fields"`

	// 分区模式。
	ShardMode *string `bson:"shard_mode,omitempty"`

	// 排序键字段名数组，顺序组合。
	SortKeyFields *[]Field `bson:"sort_key_fields,omitempty"`

	// 摘要字段名数组。
	AbstractFields *[]string `bson:"abstract_fields,omitempty"`

	ProvisionedThroughput *ProvisionedThroughput `bson:"provisioned_throughput,omitempty"`
}

func (o GlobalSecondaryIndex) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalSecondaryIndex struct{}"
	}

	return strings.Join([]string{"GlobalSecondaryIndex", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchWriteKvRequestBody struct {

	// 行操作数组，可以是多个表的操作。
	TableOpers []TableBatch `bson:"table_opers"`
}

func (o BatchWriteKvRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchWriteKvRequestBody struct{}"
	}

	return strings.Join([]string{"BatchWriteKvRequestBody", string(data)}, " ")
}

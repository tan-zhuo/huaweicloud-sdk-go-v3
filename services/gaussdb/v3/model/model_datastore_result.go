package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatastoreResult struct {

	// 数据库引擎，现在只支持gaussdb-mysql
	Type string `json:"type"`

	// 数据库版本。数据库支持的详细版本信息，可调用查询数据库引擎的版本接口获取。
	Version string `json:"version"`
}

func (o DatastoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreResult struct{}"
	}

	return strings.Join([]string{"DatastoreResult", string(data)}, " ")
}

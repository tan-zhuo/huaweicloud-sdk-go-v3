package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MysqlDatastoreWithKernelVersion 数据库信息。
type MysqlDatastoreWithKernelVersion struct {

	// 数据库引擎，现在只支持gaussdb-mysql。
	Type string `json:"type"`

	// 兼容的开源数据库版本号，返回三位开源版本号。
	Version string `json:"version"`

	// 数据库内核版本
	KernelVersion string `json:"kernel_version"`
}

func (o MysqlDatastoreWithKernelVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlDatastoreWithKernelVersion struct{}"
	}

	return strings.Join([]string{"MysqlDatastoreWithKernelVersion", string(data)}, " ")
}

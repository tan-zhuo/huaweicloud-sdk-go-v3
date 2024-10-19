package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesDatastoreResult 数据库信息。
type ListInstancesDatastoreResult struct {

	// 数据库引擎。
	Type string `json:"type"`

	// 数据库版本号。
	Version string `json:"version"`

	// 是否有补丁版本的数据库支持升级，返回true时可以通过升级补丁接口进行数据库升级，否则不允许升级补丁。
	PatchAvailable bool `json:"patch_available"`

	// 数据库的完整版本号(目前只有GeminiDB Cassandra支持)。
	WholeVersion string `json:"whole_version"`
}

func (o ListInstancesDatastoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesDatastoreResult struct{}"
	}

	return strings.Join([]string{"ListInstancesDatastoreResult", string(data)}, " ")
}

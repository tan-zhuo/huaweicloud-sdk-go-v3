package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtendReplicationRequestBody 复制对扩容请求体
type ExtendReplicationRequestBody struct {
	ExtendReplication *ExtendReplicationRequestParams `json:"extend-replication"`
}

func (o ExtendReplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendReplicationRequestBody struct{}"
	}

	return strings.Join([]string{"ExtendReplicationRequestBody", string(data)}, " ")
}

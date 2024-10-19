package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReplicationRequestBody 创建复制对请求体
type CreateReplicationRequestBody struct {
	Replication *CreateReplicationRequestParams `json:"replication"`
}

func (o CreateReplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateReplicationRequestBody", string(data)}, " ")
}

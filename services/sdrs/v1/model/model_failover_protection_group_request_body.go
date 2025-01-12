package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FailoverProtectionGroupRequestBody 保护组故障切换请求体
type FailoverProtectionGroupRequestBody struct {

	// 标识保护组故障切换操作。该参数目前默认值为空。
	FailoverServerGroup *interface{} `json:"failover-server-group"`
}

func (o FailoverProtectionGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailoverProtectionGroupRequestBody struct{}"
	}

	return strings.Join([]string{"FailoverProtectionGroupRequestBody", string(data)}, " ")
}

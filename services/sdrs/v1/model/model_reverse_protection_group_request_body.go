package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReverseProtectionGroupRequestBody 保护组切换请求体
type ReverseProtectionGroupRequestBody struct {
	ReverseServerGroup *ReverseProtectionGroupRequestParams `json:"reverse-server-group"`
}

func (o ReverseProtectionGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReverseProtectionGroupRequestBody struct{}"
	}

	return strings.Join([]string{"ReverseProtectionGroupRequestBody", string(data)}, " ")
}

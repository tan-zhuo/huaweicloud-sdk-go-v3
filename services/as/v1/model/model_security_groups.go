package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroups 安全组信息
type SecurityGroups struct {

	// 安全组ID
	Id string `json:"id"`
}

func (o SecurityGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroups struct{}"
	}

	return strings.Join([]string{"SecurityGroups", string(data)}, " ")
}

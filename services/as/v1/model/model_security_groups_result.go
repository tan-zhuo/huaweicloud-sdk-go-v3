package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroupsResult 安全组信息
type SecurityGroupsResult struct {

	// 安全组ID
	Id *string `json:"id,omitempty"`
}

func (o SecurityGroupsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupsResult struct{}"
	}

	return strings.Join([]string{"SecurityGroupsResult", string(data)}, " ")
}

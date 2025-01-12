package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityGroupRequest Request Object
type DeleteSecurityGroupRequest struct {

	// 安全组ID
	SecurityGroupId string `json:"security_group_id"`
}

func (o DeleteSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupRequest", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeleteSecurityGroupRequest Request Object
type NeutronDeleteSecurityGroupRequest struct {

	// 安全组ID
	SecurityGroupId string `json:"security_group_id"`
}

func (o NeutronDeleteSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSecurityGroupRequest", string(data)}, " ")
}

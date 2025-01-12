package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateSecurityGroupRequest Request Object
type NeutronUpdateSecurityGroupRequest struct {

	// 安全组ID
	SecurityGroupId string `json:"security_group_id"`

	Body *NeutronUpdateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSecurityGroupRequest", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityGroupRuleResponse Response Object
type ShowSecurityGroupRuleResponse struct {
	SecurityGroupRule *SecurityGroupRule `json:"security_group_rule,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o ShowSecurityGroupRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityGroupRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityGroupRuleResponse", string(data)}, " ")
}

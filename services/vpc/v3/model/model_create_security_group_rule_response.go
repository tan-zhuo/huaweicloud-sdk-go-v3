package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityGroupRuleResponse Response Object
type CreateSecurityGroupRuleResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	SecurityGroupRule *SecurityGroupRule `json:"security_group_rule,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o CreateSecurityGroupRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRuleResponse", string(data)}, " ")
}

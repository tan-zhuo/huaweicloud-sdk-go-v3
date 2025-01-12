package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityGroupRuleRequest Request Object
type CreateSecurityGroupRuleRequest struct {
	Body *CreateSecurityGroupRuleRequestBody `json:"body,omitempty"`
}

func (o CreateSecurityGroupRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRuleRequest", string(data)}, " ")
}

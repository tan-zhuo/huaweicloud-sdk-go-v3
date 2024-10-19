package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclRuleResponse Response Object
type UpdateAclRuleResponse struct {
	Data           *RuleId `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAclRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateAclRuleResponse", string(data)}, " ")
}

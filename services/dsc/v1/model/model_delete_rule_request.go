package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRuleRequest Request Object
type DeleteRuleRequest struct {

	// 规则ID
	RuleId string `json:"rule_id"`
}

func (o DeleteRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteRuleRequest", string(data)}, " ")
}

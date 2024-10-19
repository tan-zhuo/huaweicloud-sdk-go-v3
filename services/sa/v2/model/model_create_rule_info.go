package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleInfo Information of rule
type CreateRuleInfo struct {
	Rule *ConditionInfo `json:"rule,omitempty"`
}

func (o CreateRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleInfo struct{}"
	}

	return strings.Join([]string{"CreateRuleInfo", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleResponse Response Object
type CreateRuleResponse struct {
	Rule           *RuleResponse `json:"rule,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateRuleResponse", string(data)}, " ")
}

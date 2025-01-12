package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRuleDetailResponse Response Object
type ShowRuleDetailResponse struct {
	Rule           *RuleResponse `json:"rule,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowRuleDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRuleDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowRuleDetailResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListActionRuleRequest Request Object
type ListActionRuleRequest struct {
}

func (o ListActionRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActionRuleRequest struct{}"
	}

	return strings.Join([]string{"ListActionRuleRequest", string(data)}, " ")
}

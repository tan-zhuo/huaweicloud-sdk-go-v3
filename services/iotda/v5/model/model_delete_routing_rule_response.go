package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRoutingRuleResponse Response Object
type DeleteRoutingRuleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRoutingRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRoutingRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRoutingRuleResponse", string(data)}, " ")
}

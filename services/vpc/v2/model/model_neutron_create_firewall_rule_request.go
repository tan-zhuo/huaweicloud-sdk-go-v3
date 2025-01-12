package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateFirewallRuleRequest Request Object
type NeutronCreateFirewallRuleRequest struct {
	Body *NeutronCreateFirewallRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateFirewallRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallRuleRequest", string(data)}, " ")
}

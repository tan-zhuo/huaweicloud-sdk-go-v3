package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateFirewallRuleRequestBody
type NeutronCreateFirewallRuleRequestBody struct {
	FirewallRule *NeutronCreateFirewallRuleOption `json:"firewall_rule"`
}

func (o NeutronCreateFirewallRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallRuleRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallRuleRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronRemoveFirewallRuleRequestBody
type NeutronRemoveFirewallRuleRequestBody struct {

	// 功能说明：待移除的ACL规则ID
	FirewallRuleId string `json:"firewall_rule_id"`
}

func (o NeutronRemoveFirewallRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronRemoveFirewallRuleRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronRemoveFirewallRuleRequestBody", string(data)}, " ")
}

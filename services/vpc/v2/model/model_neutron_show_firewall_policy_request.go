package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowFirewallPolicyRequest Request Object
type NeutronShowFirewallPolicyRequest struct {

	// 网络ACL防火墙策略ID
	FirewallPolicyId string `json:"firewall_policy_id"`
}

func (o NeutronShowFirewallPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFirewallPolicyRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowFirewallPolicyRequest", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeleteFirewallPolicyRequest Request Object
type NeutronDeleteFirewallPolicyRequest struct {

	// 网络ACL防火墙策略ID
	FirewallPolicyId string `json:"firewall_policy_id"`
}

func (o NeutronDeleteFirewallPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFirewallPolicyRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFirewallPolicyRequest", string(data)}, " ")
}

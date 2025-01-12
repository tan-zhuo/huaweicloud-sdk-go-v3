package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowFirewallPolicyResponse Response Object
type NeutronShowFirewallPolicyResponse struct {
	FirewallPolicy *NeutronFirewallPolicy `json:"firewall_policy,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o NeutronShowFirewallPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFirewallPolicyResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowFirewallPolicyResponse", string(data)}, " ")
}

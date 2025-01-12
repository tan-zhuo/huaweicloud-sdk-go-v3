package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateFirewallPolicyResponse Response Object
type NeutronCreateFirewallPolicyResponse struct {
	FirewallPolicy *NeutronFirewallPolicy `json:"firewall_policy,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o NeutronCreateFirewallPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallPolicyResponse struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallPolicyResponse", string(data)}, " ")
}

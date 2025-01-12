package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateFirewallGroupRequestBody
type NeutronCreateFirewallGroupRequestBody struct {
	FirewallGroup *NeutronCreateFirewallGroupOption `json:"firewall_group"`
}

func (o NeutronCreateFirewallGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallGroupRequestBody", string(data)}, " ")
}

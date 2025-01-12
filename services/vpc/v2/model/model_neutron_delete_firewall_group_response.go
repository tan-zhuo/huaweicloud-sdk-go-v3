package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeleteFirewallGroupResponse Response Object
type NeutronDeleteFirewallGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteFirewallGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFirewallGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFirewallGroupResponse", string(data)}, " ")
}

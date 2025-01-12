package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNatGatewayDnatRuleRequest Request Object
type DeleteNatGatewayDnatRuleRequest struct {

	// 公网NAT网关实例的ID。
	NatGatewayId string `json:"nat_gateway_id"`

	// DNAT规则的ID。
	DnatRuleId string `json:"dnat_rule_id"`
}

func (o DeleteNatGatewayDnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayDnatRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayDnatRuleRequest", string(data)}, " ")
}

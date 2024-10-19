package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatGatewaySnatRuleOption 更新SNAT规则的请求体。
type UpdateNatGatewaySnatRuleOption struct {

	// 公网NAT网关的id。
	NatGatewayId string `json:"nat_gateway_id"`

	// 功能说明：弹性公网IP，多个弹性公网IP使用逗号分隔。 约束：弹性公网IP的id个数不能超过20个
	PublicIpAddress *string `json:"public_ip_address,omitempty"`

	// 全域弹性公网IP的id。
	GlobalEipId *string `json:"global_eip_id,omitempty"`

	// SNAT规则的描述，长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`
}

func (o UpdateNatGatewaySnatRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewaySnatRuleOption struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewaySnatRuleOption", string(data)}, " ")
}

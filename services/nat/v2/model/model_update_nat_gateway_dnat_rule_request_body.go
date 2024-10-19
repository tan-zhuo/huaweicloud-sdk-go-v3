package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatGatewayDnatRuleRequestBody 更新DNAT规则的请求体。
type UpdateNatGatewayDnatRuleRequestBody struct {
	DnatRule *UpdateNatGatewayDnatRuleOption `json:"dnat_rule"`
}

func (o UpdateNatGatewayDnatRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayDnatRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayDnatRuleRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNatGatewaySnatRuleRequest Request Object
type ShowNatGatewaySnatRuleRequest struct {

	// SNAT规则的ID。
	SnatRuleId string `json:"snat_rule_id"`
}

func (o ShowNatGatewaySnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNatGatewaySnatRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowNatGatewaySnatRuleRequest", string(data)}, " ")
}

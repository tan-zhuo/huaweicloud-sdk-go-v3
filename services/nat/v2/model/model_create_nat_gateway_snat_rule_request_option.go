package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNatGatewaySnatRuleRequestOption 创建SNAT规则的请求体。
type CreateNatGatewaySnatRuleRequestOption struct {
	SnatRule *CreateNatGatewaySnatRuleOption `json:"snat_rule"`
}

func (o CreateNatGatewaySnatRuleRequestOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewaySnatRuleRequestOption struct{}"
	}

	return strings.Join([]string{"CreateNatGatewaySnatRuleRequestOption", string(data)}, " ")
}

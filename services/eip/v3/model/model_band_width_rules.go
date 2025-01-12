package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandWidthRules 带宽规则对象
type BandWidthRules struct {

	// - 功能说明：带宽规则ID
	Id *string `json:"id,omitempty"`

	// - 功能说明：带宽规则名称
	Name *string `json:"name,omitempty"`

	// - 功能说明：配置状态，为False时配置不生效。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// - 功能说明：出网带宽最大值，单位Mbps - 取值范围[0,n]，其中n为所属带宽的带宽大小（size字段）。0表示设置为最大带宽
	EgressSize *int32 `json:"egress_size,omitempty"`

	// - 功能说明：出网保障带宽大小，单位Mbps - 取值范围[0,x]，其中x为所属带宽剩余的保障额
	EgressGuarentedSize *int32 `json:"egress_guarented_size,omitempty"`

	// - 功能说明：带宽对应的弹性公网IP信息 - 约束：WHOLE类型的带宽支持多个弹性公网IP，PER类型的带宽只能对应一个弹性公网IP
	PublicipInfo *[]PublicipInfoResponseBody `json:"publicip_info,omitempty"`
}

func (o BandWidthRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandWidthRules struct{}"
	}

	return strings.Join([]string{"BandWidthRules", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTrafficMirrorFilterRuleOption
type UpdateTrafficMirrorFilterRuleOption struct {

	// 功能说明：流量镜像筛选规则的描述信息 取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`

	// 功能说明：镜像流量的协议类型 取值范围：TCP、UDP、ICMP、ICMPV6、ALL
	Protocol *string `json:"protocol,omitempty"`

	// 功能说明：镜像流量的地址协议版本 取值范围：IPv4，IPv6
	Ethertype *string `json:"ethertype,omitempty"`

	// 功能说明：镜像流量的源网段
	SourceCidrBlock *string `json:"source_cidr_block,omitempty"`

	// 功能说明：镜像流量的目的网段
	DestinationCidrBlock *string `json:"destination_cidr_block,omitempty"`

	// 功能说明：流量源端口范围 取值范围：1~65535 格式：80-200
	SourcePortRange *string `json:"source_port_range,omitempty"`

	// 功能说明：流量目的端口范围 取值范围：1~65535 格式：80-200
	DestinationPortRange *string `json:"destination_port_range,omitempty"`

	// 功能说明：镜像规则优先级 取值范围：1~65535，数字越小，优先级越高
	Priority *int32 `json:"priority,omitempty"`

	// 功能说明：镜像策略 取值范围：accept（采集）、reject（不采集）
	Action *string `json:"action,omitempty"`
}

func (o UpdateTrafficMirrorFilterRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrafficMirrorFilterRuleOption struct{}"
	}

	return strings.Join([]string{"UpdateTrafficMirrorFilterRuleOption", string(data)}, " ")
}

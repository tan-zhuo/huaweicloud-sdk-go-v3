package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpnGatewayAvailabilityZones struct {

	// VPC关联类型的可用区列表
	Vpc *[]string `json:"vpc,omitempty"`

	// ER关联类型的可用区列表
	Er *[]string `json:"er,omitempty"`
}

func (o VpnGatewayAvailabilityZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnGatewayAvailabilityZones struct{}"
	}

	return strings.Join([]string{"VpnGatewayAvailabilityZones", string(data)}, " ")
}

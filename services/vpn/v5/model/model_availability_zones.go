package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailabilityZones struct {
	Basic *VpnGatewayAvailabilityZones `json:"basic,omitempty"`

	Professional1 *VpnGatewayAvailabilityZones `json:"professional1,omitempty"`

	Professional2 *VpnGatewayAvailabilityZones `json:"professional2,omitempty"`

	Professional1NonFixedIP *VpnGatewayAvailabilityZones `json:"Professional1-NonFixedIP,omitempty"`

	Professional2NonFixedIP *VpnGatewayAvailabilityZones `json:"Professional2-NonFixedIP,omitempty"`

	Gm *VpnGatewayAvailabilityZones `json:"gm,omitempty"`
}

func (o AvailabilityZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailabilityZones struct{}"
	}

	return strings.Join([]string{"AvailabilityZones", string(data)}, " ")
}

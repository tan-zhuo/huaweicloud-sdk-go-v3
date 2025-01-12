package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFloatingIpOption 更新floatingip对象
type UpdateFloatingIpOption struct {

	// 端口id。
	PortId *string `json:"port_id,omitempty"`
}

func (o UpdateFloatingIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFloatingIpOption struct{}"
	}

	return strings.Join([]string{"UpdateFloatingIpOption", string(data)}, " ")
}

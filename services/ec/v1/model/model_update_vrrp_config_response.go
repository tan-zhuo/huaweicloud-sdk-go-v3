package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVrrpConfigResponse Response Object
type UpdateVrrpConfigResponse struct {

	// 智能企业网关ID
	IegId *string `json:"ieg_id,omitempty"`

	// 虚路由ID
	VirtualRouterId *int32 `json:"virtual_router_id,omitempty"`

	// 虚IP
	VirtualIp *string `json:"virtual_ip,omitempty"`

	// 主设备ID
	ActiveEquipmentId *string `json:"active_equipment_id,omitempty"`

	// 主设备接口名字
	ActiveInterfaceName *string `json:"active_interface_name,omitempty"`

	// 备设备ID
	StandbyEquipmentId *string `json:"standby_equipment_id,omitempty"`

	// 备设备接口名字
	StandbyInterfaceName *string `json:"standby_interface_name,omitempty"`
	HttpStatusCode       int     `json:"-"`
}

func (o UpdateVrrpConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVrrpConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateVrrpConfigResponse", string(data)}, " ")
}

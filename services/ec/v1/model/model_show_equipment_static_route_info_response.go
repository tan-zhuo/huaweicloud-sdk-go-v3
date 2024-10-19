package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEquipmentStaticRouteInfoResponse Response Object
type ShowEquipmentStaticRouteInfoResponse struct {

	// 设备静态路由配置列表
	StaticRoutes   *[]StaticRouteItem `json:"static_routes,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowEquipmentStaticRouteInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEquipmentStaticRouteInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowEquipmentStaticRouteInfoResponse", string(data)}, " ")
}

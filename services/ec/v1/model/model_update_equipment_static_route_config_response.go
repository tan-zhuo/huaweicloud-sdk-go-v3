package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentStaticRouteConfigResponse Response Object
type UpdateEquipmentStaticRouteConfigResponse struct {

	// 目标网络
	Prefix *string `json:"prefix,omitempty"`

	// 下一跳地址
	NextHop *string `json:"next_hop,omitempty"`

	// 接口名字
	InterfaceName *string `json:"interface_name,omitempty"`

	// 优先级
	Priority *int32 `json:"priority,omitempty"`

	// 自动检测
	TrackNqa *bool `json:"track_nqa,omitempty"`

	// 发布到企业连接网络
	PostToCloud    *bool `json:"post_to_cloud,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateEquipmentStaticRouteConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentStaticRouteConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentStaticRouteConfigResponse", string(data)}, " ")
}

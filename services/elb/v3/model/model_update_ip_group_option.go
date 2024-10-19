package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateIpGroupOption struct {

	// 参数解释：IP地址组的描述信息
	Description *string `json:"description,omitempty"`

	// 参数解释：IP地址组的名称
	Name *string `json:"name,omitempty"`

	// 参数解释：IP地址组中包含的IP列表。
	IpList *[]UpadateIpGroupIpOption `json:"ip_list,omitempty"`
}

func (o UpdateIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupOption", string(data)}, " ")
}

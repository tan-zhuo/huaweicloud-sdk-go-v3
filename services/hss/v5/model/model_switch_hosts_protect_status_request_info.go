package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 切换防护的请求信息
type SwitchHostsProtectStatusRequestInfo struct {

	// 开通的版本，包含如下:   - hss.version.null : VERSION_NULL   - hss.version.basic : VERSION_BASIC   - hss.version.enterprise : VERSION_ENTERPRISE   - hss.version.premium : VERSION_PREMIUM   - hss.version.wtp : VERSION_WTP
	Version *string `json:"version,omitempty"`

	// 付费模式   - packet_cycle : 包周期   - on_demand : 按需
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 资源实例ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 服务器列表
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// 资源标签
	Tags *[]TagInfo `json:"tags,omitempty"`
}

func (o SwitchHostsProtectStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchHostsProtectStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"SwitchHostsProtectStatusRequestInfo", string(data)}, " ")
}
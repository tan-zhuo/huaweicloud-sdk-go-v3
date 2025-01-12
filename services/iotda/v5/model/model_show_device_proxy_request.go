package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceProxyRequest Request Object
type ShowDeviceProxyRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，建议携带该参数，在使用专业版时必须携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID，具体获取方式请参考[[查看实例详情](https://support.huaweicloud.com/usermanual-iothub/iot_01_0121.html)](tag:hws) [[查看实例详情](https://support.huaweicloud.com/intl/zh-cn/usermanual-iothub/iot_01_0121.html)](tag:hws_hk)。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：设备代理ID，用于唯一标识一个设备代理。在注册设备代理时由物联网平台分配获得。
	ProxyId string `json:"proxy_id"`
}

func (o ShowDeviceProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceProxyRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceProxyRequest", string(data)}, " ")
}

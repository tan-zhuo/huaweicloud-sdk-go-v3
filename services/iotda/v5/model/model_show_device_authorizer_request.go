package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceAuthorizerRequest Request Object
type ShowDeviceAuthorizerRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，建议携带该参数，在使用专业版时必须携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID，具体获取方式请参考[[查看实例详情](https://support.huaweicloud.com/usermanual-iothub/iot_01_0079.html#section1)](tag:hws) [[查看实例详情](https://support.huaweicloud.com/intl/zh-cn/usermanual-iothub/iot_01_0079.html#section1)](tag:hws_hk)。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// 自定义鉴权ID
	AuthorizerId string `json:"authorizer_id"`
}

func (o ShowDeviceAuthorizerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceAuthorizerRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceAuthorizerRequest", string(data)}, " ")
}

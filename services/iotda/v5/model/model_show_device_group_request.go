package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceGroupRequest Request Object
type ShowDeviceGroupRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：设备组ID，用于唯一标识一个设备组，在创建设备组时由物联网平台分配。 **取值范围**：长度不超过36，十六进制字符串和连接符（-）的组合。
	GroupId string `json:"group_id"`
}

func (o ShowDeviceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceGroupRequest", string(data)}, " ")
}

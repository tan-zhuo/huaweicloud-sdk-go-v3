package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOtaPackageRequest Request Object
type DeleteOtaPackageRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：升级包ID，用于唯一标识一个升级包。由物联网平台分配获得。 **取值范围**：长度不超过36，只允许字母、数字、连接符（-）的组合。
	PackageId string `json:"package_id"`
}

func (o DeleteOtaPackageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOtaPackageRequest struct{}"
	}

	return strings.Join([]string{"DeleteOtaPackageRequest", string(data)}, " ")
}

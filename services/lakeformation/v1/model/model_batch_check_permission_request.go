package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckPermissionRequest Request Object
type BatchCheckPermissionRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	Body *CheckPermissionInput `json:"body,omitempty"`
}

func (o BatchCheckPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckPermissionRequest struct{}"
	}

	return strings.Join([]string{"BatchCheckPermissionRequest", string(data)}, " ")
}

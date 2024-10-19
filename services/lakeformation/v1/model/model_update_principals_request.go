package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrincipalsRequest Request Object
type UpdatePrincipalsRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// 角色名称。只能包含字母、数字和下划线，且长度为1~255个字符。
	RoleName string `json:"role_name"`

	Body *[]Principal `json:"body,omitempty"`
}

func (o UpdatePrincipalsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrincipalsRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrincipalsRequest", string(data)}, " ")
}

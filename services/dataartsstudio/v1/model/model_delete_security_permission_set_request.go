package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityPermissionSetRequest Request Object
type DeleteSecurityPermissionSetRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o DeleteSecurityPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityPermissionSetRequest", string(data)}, " ")
}

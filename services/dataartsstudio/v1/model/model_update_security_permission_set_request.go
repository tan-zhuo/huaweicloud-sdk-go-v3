package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityPermissionSetRequest Request Object
type UpdateSecurityPermissionSetRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *PermissionSetCreateDto `json:"body,omitempty"`
}

func (o UpdateSecurityPermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPermissionSetRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPermissionSetRequest", string(data)}, " ")
}

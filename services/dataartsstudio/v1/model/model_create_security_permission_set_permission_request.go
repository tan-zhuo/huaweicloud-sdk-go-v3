package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityPermissionSetPermissionRequest Request Object
type CreateSecurityPermissionSetPermissionRequest struct {

	// 权限集id
	PermissionSetId string `json:"permission_set_id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *PermissionSetPermissionCreateDto `json:"body,omitempty"`
}

func (o CreateSecurityPermissionSetPermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPermissionSetPermissionRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityPermissionSetPermissionRequest", string(data)}, " ")
}

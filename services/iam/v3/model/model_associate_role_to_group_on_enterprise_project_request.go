package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateRoleToGroupOnEnterpriseProjectRequest Request Object
type AssociateRoleToGroupOnEnterpriseProjectRequest struct {

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 用户组ID。
	GroupId string `json:"group_id"`

	// 权限ID。
	RoleId string `json:"role_id"`
}

func (o AssociateRoleToGroupOnEnterpriseProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRoleToGroupOnEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"AssociateRoleToGroupOnEnterpriseProjectRequest", string(data)}, " ")
}

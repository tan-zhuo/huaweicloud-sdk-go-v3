package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRolesForUserOnEnterpriseProjectResponse Response Object
type ListRolesForUserOnEnterpriseProjectResponse struct {

	// 角色列表。
	Roles          *[]RolesItem `json:"roles,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListRolesForUserOnEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRolesForUserOnEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"ListRolesForUserOnEnterpriseProjectResponse", string(data)}, " ")
}

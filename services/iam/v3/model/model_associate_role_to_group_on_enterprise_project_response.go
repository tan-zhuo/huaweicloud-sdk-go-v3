package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateRoleToGroupOnEnterpriseProjectResponse Response Object
type AssociateRoleToGroupOnEnterpriseProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateRoleToGroupOnEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRoleToGroupOnEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"AssociateRoleToGroupOnEnterpriseProjectResponse", string(data)}, " ")
}

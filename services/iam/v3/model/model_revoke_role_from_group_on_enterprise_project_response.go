package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokeRoleFromGroupOnEnterpriseProjectResponse Response Object
type RevokeRoleFromGroupOnEnterpriseProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RevokeRoleFromGroupOnEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeRoleFromGroupOnEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"RevokeRoleFromGroupOnEnterpriseProjectResponse", string(data)}, " ")
}

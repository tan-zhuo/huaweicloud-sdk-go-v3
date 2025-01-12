package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokeRoleFromUserOnEnterpriseProjectResponse Response Object
type RevokeRoleFromUserOnEnterpriseProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RevokeRoleFromUserOnEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeRoleFromUserOnEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"RevokeRoleFromUserOnEnterpriseProjectResponse", string(data)}, " ")
}

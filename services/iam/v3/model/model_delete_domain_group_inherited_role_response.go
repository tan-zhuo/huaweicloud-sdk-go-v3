package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainGroupInheritedRoleResponse Response Object
type DeleteDomainGroupInheritedRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDomainGroupInheritedRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainGroupInheritedRoleResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainGroupInheritedRoleResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListProjectPermissionsForGroupResponse Response Object
type KeystoneListProjectPermissionsForGroupResponse struct {
	Links *Links `json:"links,omitempty"`

	// 权限信息列表。
	Roles          *[]RoleResult `json:"roles,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o KeystoneListProjectPermissionsForGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListProjectPermissionsForGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListProjectPermissionsForGroupResponse", string(data)}, " ")
}

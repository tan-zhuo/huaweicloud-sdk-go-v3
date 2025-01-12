package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListPermissionsResponse Response Object
type KeystoneListPermissionsResponse struct {
	Links *Links `json:"links,omitempty"`

	// 权限信息列表。
	Roles *[]RoleResult `json:"roles,omitempty"`

	// 在查询参数存在domain_id时，返回自定义策略总数
	TotalNumber    *int32 `json:"total_number,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o KeystoneListPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListPermissionsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListPermissionsResponse", string(data)}, " ")
}

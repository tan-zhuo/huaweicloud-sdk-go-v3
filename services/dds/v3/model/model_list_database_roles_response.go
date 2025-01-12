package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseRolesResponse Response Object
type ListDatabaseRolesResponse struct {

	// 数据库角色信息。
	Roles *string `json:"roles,omitempty"`

	// 数据库角色总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabaseRolesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseRolesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseRolesResponse", string(data)}, " ")
}

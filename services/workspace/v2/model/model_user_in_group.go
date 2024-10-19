package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserInGroup 查询桌面用户组中的桌面用户响应的用户信息。
type UserInGroup struct {

	// 用户ID。
	Id *string `json:"id,omitempty"`

	// 桌面用户名。
	UserName *string `json:"user_name,omitempty"`

	// 用户邮箱。
	UserEmail *string `json:"user_email,omitempty"`

	// 用户手机号。
	UserPhone *string `json:"user_phone,omitempty"`

	// 用户描述。
	Description *string `json:"description,omitempty"`
}

func (o UserInGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserInGroup struct{}"
	}

	return strings.Join([]string{"UserInGroup", string(data)}, " ")
}

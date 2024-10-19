package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserGroupResponse Response Object
type CreateUserGroupResponse struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUserGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateUserGroupResponse", string(data)}, " ")
}

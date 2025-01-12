package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueUser struct {

	// 用户uuid
	UserId *string `json:"user_id,omitempty"`

	// 用户数字id
	UserNumId *int32 `json:"user_num_id,omitempty"`

	// 用户id
	Id *int32 `json:"id,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`
}

func (o IssueUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueUser struct{}"
	}

	return strings.Join([]string{"IssueUser", string(data)}, " ")
}

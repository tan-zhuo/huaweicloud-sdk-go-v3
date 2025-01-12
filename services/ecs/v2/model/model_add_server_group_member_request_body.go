package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServerGroupMemberRequestBody This is a auto create Body Object
type AddServerGroupMemberRequestBody struct {
	AddMember *ServerGroupMember `json:"add_member"`
}

func (o AddServerGroupMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServerGroupMemberRequestBody struct{}"
	}

	return strings.Join([]string{"AddServerGroupMemberRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerGroupMemberRequestBody This is a auto create Body Object
type DeleteServerGroupMemberRequestBody struct {
	RemoveMember *ServerGroupMember `json:"remove_member"`
}

func (o DeleteServerGroupMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupMemberRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupMemberRequestBody", string(data)}, " ")
}

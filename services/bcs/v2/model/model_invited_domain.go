package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvitedDomain 被邀请列表
type InvitedDomain struct {

	// 被邀请方租户
	InvitedUser string `json:"invited_user"`
}

func (o InvitedDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvitedDomain struct{}"
	}

	return strings.Join([]string{"InvitedDomain", string(data)}, " ")
}

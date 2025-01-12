package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListGroupsForUserResponse Response Object
type KeystoneListGroupsForUserResponse struct {

	// 用户组信息列表。
	Groups *[]KeystoneGroupResult `json:"groups,omitempty"`

	Links          *Links `json:"links,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o KeystoneListGroupsForUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListGroupsForUserResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListGroupsForUserResponse", string(data)}, " ")
}

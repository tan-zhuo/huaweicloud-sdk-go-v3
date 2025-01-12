package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListUsersResponse Response Object
type KeystoneListUsersResponse struct {
	Links *Links `json:"links,omitempty"`

	// IAM用户信息列表。
	Users          *[]KeystoneListUsersResult `json:"users,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o KeystoneListUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListUsersResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListUsersResponse", string(data)}, " ")
}

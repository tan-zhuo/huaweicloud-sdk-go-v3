package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCheckUserInGroupResponse Response Object
type KeystoneCheckUserInGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneCheckUserInGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCheckUserInGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCheckUserInGroupResponse", string(data)}, " ")
}

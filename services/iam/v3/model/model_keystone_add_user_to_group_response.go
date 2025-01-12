package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneAddUserToGroupResponse Response Object
type KeystoneAddUserToGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneAddUserToGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneAddUserToGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneAddUserToGroupResponse", string(data)}, " ")
}

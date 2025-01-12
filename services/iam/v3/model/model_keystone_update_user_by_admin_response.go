package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateUserByAdminResponse Response Object
type KeystoneUpdateUserByAdminResponse struct {
	User           *KeystoneUpdateUserByAdminResult `json:"user,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o KeystoneUpdateUserByAdminResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateUserByAdminResponse struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateUserByAdminResponse", string(data)}, " ")
}

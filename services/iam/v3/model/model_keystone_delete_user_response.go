package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneDeleteUserResponse Response Object
type KeystoneDeleteUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneDeleteUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneDeleteUserResponse struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteUserResponse", string(data)}, " ")
}

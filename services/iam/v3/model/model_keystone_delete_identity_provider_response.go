package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneDeleteIdentityProviderResponse Response Object
type KeystoneDeleteIdentityProviderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneDeleteIdentityProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneDeleteIdentityProviderResponse struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteIdentityProviderResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateIdentityProviderRequestBody
type KeystoneUpdateIdentityProviderRequestBody struct {
	IdentityProvider *UpdateIdentityproviderOption `json:"identity_provider"`
}

func (o KeystoneUpdateIdentityProviderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateIdentityProviderRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateIdentityProviderRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateIdentityProviderResponse Response Object
type KeystoneUpdateIdentityProviderResponse struct {
	IdentityProvider *IdentityprovidersResult `json:"identity_provider,omitempty"`
	HttpStatusCode   int                      `json:"-"`
}

func (o KeystoneUpdateIdentityProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateIdentityProviderResponse struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateIdentityProviderResponse", string(data)}, " ")
}

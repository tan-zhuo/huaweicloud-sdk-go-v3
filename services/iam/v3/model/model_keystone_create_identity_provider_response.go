package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateIdentityProviderResponse Response Object
type KeystoneCreateIdentityProviderResponse struct {
	IdentityProvider *IdentityprovidersResult `json:"identity_provider,omitempty"`
	HttpStatusCode   int                      `json:"-"`
}

func (o KeystoneCreateIdentityProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateIdentityProviderResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateIdentityProviderResponse", string(data)}, " ")
}

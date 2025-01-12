package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowIdentityProviderResponse Response Object
type KeystoneShowIdentityProviderResponse struct {
	IdentityProvider *IdentityprovidersResult `json:"identity_provider,omitempty"`
	HttpStatusCode   int                      `json:"-"`
}

func (o KeystoneShowIdentityProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowIdentityProviderResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowIdentityProviderResponse", string(data)}, " ")
}

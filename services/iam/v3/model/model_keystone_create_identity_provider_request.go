package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateIdentityProviderRequest Request Object
type KeystoneCreateIdentityProviderRequest struct {

	// 待注册的身份提供商ID。
	Id string `json:"id"`

	Body *KeystoneCreateIdentityProviderRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateIdentityProviderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateIdentityProviderRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCreateIdentityProviderRequest", string(data)}, " ")
}

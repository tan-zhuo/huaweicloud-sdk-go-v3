package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateScopedTokenRequestBody
type KeystoneCreateScopedTokenRequestBody struct {
	Auth *ScopedTokenAuth `json:"auth"`
}

func (o KeystoneCreateScopedTokenRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateScopedTokenRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateScopedTokenRequestBody", string(data)}, " ")
}

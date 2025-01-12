package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateUserTokenByPasswordRequestBody
type KeystoneCreateUserTokenByPasswordRequestBody struct {
	Auth *PwdAuth `json:"auth"`
}

func (o KeystoneCreateUserTokenByPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserTokenByPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordRequestBody", string(data)}, " ")
}

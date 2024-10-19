package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneDeleteProtocolResponse Response Object
type KeystoneDeleteProtocolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneDeleteProtocolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneDeleteProtocolResponse struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteProtocolResponse", string(data)}, " ")
}

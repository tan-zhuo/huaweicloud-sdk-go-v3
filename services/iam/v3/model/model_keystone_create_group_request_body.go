package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateGroupRequestBody
type KeystoneCreateGroupRequestBody struct {
	Group *KeystoneCreateGroupOption `json:"group"`
}

func (o KeystoneCreateGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateGroupRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateGroupRequestBody", string(data)}, " ")
}

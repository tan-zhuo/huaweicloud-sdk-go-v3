package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerPasswordResponse Response Object
type ShowServerPasswordResponse struct {

	// 加密后的密码。
	Password       *string `json:"password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowServerPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerPasswordResponse struct{}"
	}

	return strings.Join([]string{"ShowServerPasswordResponse", string(data)}, " ")
}

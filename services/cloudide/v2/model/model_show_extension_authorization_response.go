package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExtensionAuthorizationResponse Response Object
type ShowExtensionAuthorizationResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowExtensionAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtensionAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"ShowExtensionAuthorizationResponse", string(data)}, " ")
}

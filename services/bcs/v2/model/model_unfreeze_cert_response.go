package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnfreezeCertResponse Response Object
type UnfreezeCertResponse struct {

	// 操作结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnfreezeCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeCertResponse struct{}"
	}

	return strings.Join([]string{"UnfreezeCertResponse", string(data)}, " ")
}

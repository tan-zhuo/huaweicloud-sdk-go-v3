package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorErrorResponse struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误消息。
	ErrorMsg string `json:"error_msg"`
}

func (o FlavorErrorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorErrorResponse struct{}"
	}

	return strings.Join([]string{"FlavorErrorResponse", string(data)}, " ")
}

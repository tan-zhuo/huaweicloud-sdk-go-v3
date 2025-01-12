package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcResponse Response Object
type UpdateVpcResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 错误消息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 错误码
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpcResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertResponse Response Object
type DeleteAlertResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *BatchOperateAlertResult `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAlertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlertResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookResponse Response Object
type DeletePlaybookResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *PlaybookInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePlaybookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookResponse struct{}"
	}

	return strings.Join([]string{"DeletePlaybookResponse", string(data)}, " ")
}

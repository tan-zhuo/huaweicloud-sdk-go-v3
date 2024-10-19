package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookRuleResponse Response Object
type DeletePlaybookRuleResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应消息
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePlaybookRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookRuleResponse struct{}"
	}

	return strings.Join([]string{"DeletePlaybookRuleResponse", string(data)}, " ")
}

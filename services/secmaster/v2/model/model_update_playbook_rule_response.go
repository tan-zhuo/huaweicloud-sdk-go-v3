package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePlaybookRuleResponse Response Object
type UpdatePlaybookRuleResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *RuleInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePlaybookRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePlaybookRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdatePlaybookRuleResponse", string(data)}, " ")
}

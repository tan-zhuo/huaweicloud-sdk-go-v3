package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRuleResponse Response Object
type DeleteRuleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRuleResponse", string(data)}, " ")
}

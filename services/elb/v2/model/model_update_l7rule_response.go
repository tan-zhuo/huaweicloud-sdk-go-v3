package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateL7ruleResponse Response Object
type UpdateL7ruleResponse struct {
	Rule           *L7ruleResp `json:"rule,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateL7ruleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7ruleResponse struct{}"
	}

	return strings.Join([]string{"UpdateL7ruleResponse", string(data)}, " ")
}

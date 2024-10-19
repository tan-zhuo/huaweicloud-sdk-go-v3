package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartRuleResponse Response Object
type StartRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartRuleResponse struct{}"
	}

	return strings.Join([]string{"StartRuleResponse", string(data)}, " ")
}

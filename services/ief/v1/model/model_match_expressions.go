package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MatchExpressions 匹配规则表达式
type MatchExpressions struct {

	// 匹配规则表达式
	MatchExpressions *[]MatchExpression `json:"matchExpressions,omitempty"`
}

func (o MatchExpressions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MatchExpressions struct{}"
	}

	return strings.Join([]string{"MatchExpressions", string(data)}, " ")
}

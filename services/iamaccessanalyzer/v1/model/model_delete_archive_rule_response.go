package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteArchiveRuleResponse Response Object
type DeleteArchiveRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteArchiveRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteArchiveRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteArchiveRuleResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecordRuleRequest Request Object
type DeleteRecordRuleRequest struct {

	// 规则ID
	Id string `json:"id"`
}

func (o DeleteRecordRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordRuleRequest", string(data)}, " ")
}

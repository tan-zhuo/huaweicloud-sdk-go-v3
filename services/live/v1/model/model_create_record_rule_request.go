package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecordRuleRequest Request Object
type CreateRecordRuleRequest struct {
	Body *RecordRuleRequest `json:"body,omitempty"`
}

func (o CreateRecordRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordRuleRequest", string(data)}, " ")
}

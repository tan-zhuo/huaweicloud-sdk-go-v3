package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecordingRuleResponse Response Object
type CreateRecordingRuleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRecordingRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordingRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateRecordingRuleResponse", string(data)}, " ")
}

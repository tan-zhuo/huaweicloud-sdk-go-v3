package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendAimBatchDifferentMessagesRequest Request Object
type SendAimBatchDifferentMessagesRequest struct {
	Body *SmsMultiTemplateTaskRequestBody `json:"body,omitempty"`
}

func (o SendAimBatchDifferentMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendAimBatchDifferentMessagesRequest struct{}"
	}

	return strings.Join([]string{"SendAimBatchDifferentMessagesRequest", string(data)}, " ")
}

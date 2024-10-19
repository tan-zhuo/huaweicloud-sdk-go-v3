package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEventRequest Request Object
type DeleteEventRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`

	// 测试事件ID
	EventId string `json:"event_id"`
}

func (o DeleteEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventRequest struct{}"
	}

	return strings.Join([]string{"DeleteEventRequest", string(data)}, " ")
}

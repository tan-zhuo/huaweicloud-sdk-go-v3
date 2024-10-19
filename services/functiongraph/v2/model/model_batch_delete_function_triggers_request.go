package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteFunctionTriggersRequest Request Object
type BatchDeleteFunctionTriggersRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o BatchDeleteFunctionTriggersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFunctionTriggersRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteFunctionTriggersRequest", string(data)}, " ")
}

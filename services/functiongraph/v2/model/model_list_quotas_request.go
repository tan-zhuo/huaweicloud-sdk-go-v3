package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotasRequest Request Object
type ListQuotasRequest struct {

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListQuotasRequest", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartInstanceRequestBody 批量启动边缘实例请求体。
type BatchStartInstanceRequestBody struct {
	OsStart *BatchStart `json:"os-start,omitempty"`
}

func (o BatchStartInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"BatchStartInstanceRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateMessageTemplateRequestBody struct {

	// 模板内容。
	Content string `json:"content"`
}

func (o UpdateMessageTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateMessageTemplateRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMessageTemplateRequest Request Object
type DeleteMessageTemplateRequest struct {

	// 模板唯一的资源标识，可通过查询[消息模板列表](ListMessageTemplates.xml)获取该标识。
	MessageTemplateId string `json:"message_template_id"`
}

func (o DeleteMessageTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMessageTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteMessageTemplateRequest", string(data)}, " ")
}

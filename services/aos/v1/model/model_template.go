package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Template struct {

	// 模板的唯一ID，由模板服务随机生成
	TemplateId string `json:"template_id"`

	// 用户希望创建的模板名称
	TemplateName string `json:"template_name"`

	// 模板的描述。可用于客户识别自己的模板
	TemplateDescription *string `json:"template_description,omitempty"`

	// 模板的生成时间，格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	CreateTime string `json:"create_time"`

	// 模板的更新时间，格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	UpdateTime string `json:"update_time"`

	// 模板中最新的模板版本ID
	LatestVersionId string `json:"latest_version_id"`

	// 模板中最新模板版本的版本描述
	LatestVersionDescription string `json:"latest_version_description"`
}

func (o Template) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Template struct{}"
	}

	return strings.Join([]string{"Template", string(data)}, " ")
}

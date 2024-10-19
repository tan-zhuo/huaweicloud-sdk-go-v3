package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateMetadataRequest Request Object
type ShowTemplateMetadataRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 用户希望创建的模板名称
	TemplateName string `json:"template_name"`

	// 模板的ID。当template_id存在时，模板服务会检查template_id是否和template_name匹配，不匹配会返回400
	TemplateId *string `json:"template_id,omitempty"`
}

func (o ShowTemplateMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateMetadataRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateMetadataRequest", string(data)}, " ")
}

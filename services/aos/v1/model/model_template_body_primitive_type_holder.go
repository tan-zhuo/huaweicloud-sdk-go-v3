package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateBodyPrimitiveTypeHolder struct {

	// HCL模板，描述了资源的目标状态。资源编排服务将比较此模板与当前远程资源的状态之间的区别。  template_body和template_uri 必须有且只有一个存在  **注意：**   * 资源栈集不支持敏感数据加密，资源编排服务会直接明文使用、log、展示、存储对应的template_body
	TemplateBody *string `json:"template_body,omitempty"`
}

func (o TemplateBodyPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateBodyPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"TemplateBodyPrimitiveTypeHolder", string(data)}, " ")
}

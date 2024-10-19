package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigurationsRequest Request Object
type UpdateInstanceConfigurationsRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceConfigurationsRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationsRequest", string(data)}, " ")
}

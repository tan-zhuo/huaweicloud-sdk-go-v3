package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShrinkGaussMySqlProxyRequest Request Object
type ShrinkGaussMySqlProxyRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 数据库代理ID，严格匹配UUID规则。
	ProxyId string `json:"proxy_id"`

	Body *ShrinkGaussMySqlProxyRequestBody `json:"body,omitempty"`
}

func (o ShrinkGaussMySqlProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkGaussMySqlProxyRequest struct{}"
	}

	return strings.Join([]string{"ShrinkGaussMySqlProxyRequest", string(data)}, " ")
}

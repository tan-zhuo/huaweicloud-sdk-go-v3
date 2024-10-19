package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRouteToEnhancedConnectionRequestBody struct {

	// 路由名称，长度限制：1-64个字符。
	Name string `json:"name"`

	// 路由网段范围。
	Cidr string `json:"cidr"`
}

func (o CreateRouteToEnhancedConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteToEnhancedConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRouteToEnhancedConnectionRequestBody", string(data)}, " ")
}

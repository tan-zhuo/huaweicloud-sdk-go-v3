package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllQuotas 配额列表
type AllQuotas struct {

	// 配额详情资源列表。
	Resources *[]AllResources `json:"resources,omitempty"`
}

func (o AllQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllQuotas struct{}"
	}

	return strings.Join([]string{"AllQuotas", string(data)}, " ")
}

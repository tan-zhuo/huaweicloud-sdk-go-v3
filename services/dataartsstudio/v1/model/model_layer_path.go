package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LayerPath struct {

	// 目录编号
	CatalogId *string `json:"catalog_id,omitempty"`

	// 路径名
	Name *string `json:"name,omitempty"`

	// 路径层级
	Order *int32 `json:"order,omitempty"`
}

func (o LayerPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayerPath struct{}"
	}

	return strings.Join([]string{"LayerPath", string(data)}, " ")
}

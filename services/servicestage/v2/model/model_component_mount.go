package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentMount struct {

	// 挂载路径
	Path string `json:"path"`

	// 挂载路径的子路径
	SubPath *string `json:"subPath,omitempty"`

	// 是否只读
	ReadOnly bool `json:"readOnly"`
}

func (o ComponentMount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentMount struct{}"
	}

	return strings.Join([]string{"ComponentMount", string(data)}, " ")
}

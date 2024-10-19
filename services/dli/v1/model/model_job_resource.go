package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobResource 创建会话请求参数resources的元素。
type JobResource struct {

	// 资源名称。
	Name *string `json:"name,omitempty"`

	// 资源类型。
	Type *string `json:"type,omitempty"`
}

func (o JobResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResource struct{}"
	}

	return strings.Join([]string{"JobResource", string(data)}, " ")
}

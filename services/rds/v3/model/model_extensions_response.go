package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionsResponse struct {

	// 插件名称。
	Name *string `json:"name,omitempty"`

	// 数据库名称。
	DatabaseName *string `json:"database_name,omitempty"`

	// 插件版本。
	Version *string `json:"version,omitempty"`

	// 可更新插件版本
	VersionUpdate *string `json:"version_update,omitempty"`

	// 依赖预加载库。
	SharedPreloadLibraries *string `json:"shared_preload_libraries,omitempty"`

	// 是否创建。
	Created *bool `json:"created,omitempty"`

	// 插件描述。
	Description *string `json:"description,omitempty"`
}

func (o ExtensionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionsResponse struct{}"
	}

	return strings.Join([]string{"ExtensionsResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginDto struct {

	// 唯一ID
	UniqueId *string `json:"unique_id,omitempty"`

	// 图标URL
	IconUrl *string `json:"icon_url,omitempty"`

	// 运行属性
	RuntimeAttribution string `json:"runtime_attribution"`

	// 插件名
	PluginName string `json:"plugin_name"`

	// 展示名
	DisplayName string `json:"display_name"`

	// 业务类型
	BusinessType string `json:"business_type"`

	// 业务类型展示名
	BusinessTypeDisplayName string `json:"business_type_display_name"`

	// 描述
	Description string `json:"description"`

	// 是否私有
	IsPrivate *int32 `json:"is_private,omitempty"`

	// 局点
	Region *string `json:"region,omitempty"`

	// 维护者
	Maintainers *string `json:"maintainers,omitempty"`

	// 版本号
	Version string `json:"version"`

	// 版本号说明
	VersionDescription *string `json:"version_description,omitempty"`

	ExecutionInfo *PluginDtoExecutionInfo `json:"execution_info"`

	// 输入信息
	InputInfo *[]PluginDtoInputInfo `json:"input_info,omitempty"`
}

func (o PluginDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginDto struct{}"
	}

	return strings.Join([]string{"PluginDto", string(data)}, " ")
}

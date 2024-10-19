package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoBusinessTypeDefinitionVoData struct {

	// 业务类型
	BusinessType *string `json:"business_type,omitempty"`

	// 展示名
	DisplayName *string `json:"display_name,omitempty"`

	// 唯一ID
	UniqueId *string `json:"unique_id,omitempty"`

	// 可编辑
	Editable *bool `json:"editable,omitempty"`

	// 可移除
	Removable *bool `json:"removable,omitempty"`

	// 可复制
	Cloneable *bool `json:"cloneable,omitempty"`

	// 禁用
	Disabled *bool `json:"disabled,omitempty"`

	// 条件
	Conditions *[]string `json:"conditions,omitempty"`

	// 插件列表
	PluginsList *[]PageInfoBusinessTypeDefinitionVoPluginsList `json:"plugins_list,omitempty"`
}

func (o PageInfoBusinessTypeDefinitionVoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoBusinessTypeDefinitionVoData struct{}"
	}

	return strings.Join([]string{"PageInfoBusinessTypeDefinitionVoData", string(data)}, " ")
}

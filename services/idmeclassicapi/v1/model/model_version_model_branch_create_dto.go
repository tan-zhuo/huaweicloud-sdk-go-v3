package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelBranchCreateDto struct {

	// 创建时间。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 扩展类型。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *ObjectReferenceParamDto `json:"tenant,omitempty"`
}

func (o VersionModelBranchCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelBranchCreateDto struct{}"
	}

	return strings.Join([]string{"VersionModelBranchCreateDto", string(data)}, " ")
}

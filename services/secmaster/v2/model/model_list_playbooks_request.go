package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybooksRequest Request Object
type ListPlaybooksRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 搜索关键字
	SearchTxt *string `json:"search_txt,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 分页查询参数。用于指定查询结果的起始位置，从0开始
	Offset int32 `json:"offset"`

	// 分页查询参数，用于指定一次查询最多的结果数，从1开始
	Limit int32 `json:"limit"`

	// 剧本描述
	Description *string `json:"description,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 剧本名称
	Name *string `json:"name,omitempty"`
}

func (o ListPlaybooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybooksRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybooksRequest", string(data)}, " ")
}

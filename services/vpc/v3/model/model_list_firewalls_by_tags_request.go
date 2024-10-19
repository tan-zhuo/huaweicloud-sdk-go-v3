package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirewallsByTagsRequest Request Object
type ListFirewallsByTagsRequest struct {

	// 功能说明：查询记录数 取值范围：1-1000 约束：默认为1000
	Limit *int32 `json:"limit,omitempty"`

	// 功能说明：索引位置， 从第一条数据偏移offset条数据后开始查询 约束：默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数
	Offset *int32 `json:"offset,omitempty"`

	Body *ListFirewallsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListFirewallsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListFirewallsByTagsRequest", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecycleInstancesRequest Request Object
type ListRecycleInstancesRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的数量，默认是100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRecycleInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListRecycleInstancesRequest", string(data)}, " ")
}

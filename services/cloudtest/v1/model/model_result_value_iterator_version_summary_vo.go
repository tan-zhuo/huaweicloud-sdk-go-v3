package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueIteratorVersionSummaryVo 请求的返回的数据对象
type ResultValueIteratorVersionSummaryVo struct {

	// 起始记录数 大于 实际总条数时， 值为0， 分页请求才有此值
	Total *int32 `json:"total,omitempty"`

	Value *IteratorVersionSummaryVo `json:"value,omitempty"`

	// 业务失败的提示内容，对内接口才有此值
	Reason *string `json:"reason,omitempty"`

	PageSize *int32 `json:"page_size,omitempty"`

	PageNo *int32 `json:"page_no,omitempty"`

	HasMore *bool `json:"has_more,omitempty"`
}

func (o ResultValueIteratorVersionSummaryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueIteratorVersionSummaryVo struct{}"
	}

	return strings.Join([]string{"ResultValueIteratorVersionSummaryVo", string(data)}, " ")
}

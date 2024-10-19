package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueQueryTaskTestCasesVo 请求的返回的数据对象
type ResultValueQueryTaskTestCasesVo struct {

	// 起始记录数 大于 实际总条数时， 值为0， 分页请求才有此值
	Total *int32 `json:"total,omitempty"`

	Value *QueryTaskTestCasesVo `json:"value,omitempty"`

	// 业务失败的提示内容，对内接口才有此值
	Reason *string `json:"reason,omitempty"`

	PageSize *int32 `json:"page_size,omitempty"`

	PageNo *int32 `json:"page_no,omitempty"`

	HasMore *bool `json:"has_more,omitempty"`
}

func (o ResultValueQueryTaskTestCasesVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueQueryTaskTestCasesVo struct{}"
	}

	return strings.Join([]string{"ResultValueQueryTaskTestCasesVo", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScattersUsingResponse Response Object
type ListScattersUsingResponse struct {

	// 返回结果
	List *[]TaskCaseResponseTimeDetailVo `json:"list,omitempty"`

	// 页码
	PageNum *int32 `json:"page_num,omitempty"`

	// 分页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 总页数
	TotalPage *int32 `json:"total_page,omitempty"`

	// 总条数
	TotalSize      *int64 `json:"total_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListScattersUsingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScattersUsingResponse struct{}"
	}

	return strings.Join([]string{"ListScattersUsingResponse", string(data)}, " ")
}

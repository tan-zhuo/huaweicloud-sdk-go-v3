package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionsResponse Response Object
type ListFunctionsResponse struct {

	// 函数列表。
	Functions *[]ListFunctionResult `json:"functions,omitempty"`

	// 函数下次记录读取位置。
	NextMarker *int32 `json:"next_marker,omitempty"`

	// 满足查询条件的函数总数。
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFunctionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionsResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionsResponse", string(data)}, " ")
}

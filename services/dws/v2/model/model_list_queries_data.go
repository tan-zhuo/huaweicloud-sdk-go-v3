package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQueriesData struct {

	// 查询数据列表。
	Queries *[]ListQueriesDto `json:"queries,omitempty"`

	Status *ListQueriesStatus `json:"status,omitempty"`
}

func (o ListQueriesData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesData struct{}"
	}

	return strings.Join([]string{"ListQueriesData", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalEipsResponseBodyPageInfo 分页页码信息
type ListGlobalEipsResponseBodyPageInfo struct {

	// 翻页时，作为后一页的marker取值
	NextMarker *string `json:"next_marker,omitempty"`

	// 翻页时，作为前一页的marker取值
	PreviousMarker *string `json:"previous_marker,omitempty"`

	// 当前页的数据总数
	CurrentCount *int32 `json:"current_count,omitempty"`
}

func (o ListGlobalEipsResponseBodyPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipsResponseBodyPageInfo struct{}"
	}

	return strings.Join([]string{"ListGlobalEipsResponseBodyPageInfo", string(data)}, " ")
}

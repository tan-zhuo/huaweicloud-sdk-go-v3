package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowlogResponse Response Object
type ListSlowlogResponse struct {

	// 慢日志总条数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 慢日志列表
	Slowlogs       *[]SlowlogItem `json:"slowlogs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSlowlogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowlogResponse struct{}"
	}

	return strings.Join([]string{"ListSlowlogResponse", string(data)}, " ")
}

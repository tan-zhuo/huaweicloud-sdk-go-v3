package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubsetsResponse Response Object
type ListSubsetsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty"`

	// 设备ID列表
	Items          *[]SubDevice `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListSubsetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubsetsResponse struct{}"
	}

	return strings.Join([]string{"ListSubsetsResponse", string(data)}, " ")
}

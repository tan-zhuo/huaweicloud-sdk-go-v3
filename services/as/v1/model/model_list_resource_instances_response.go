package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceInstancesResponse Response Object
type ListResourceInstancesResponse struct {

	// 标签资源实例。
	Resources *[]Resources `json:"resources,omitempty"`

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 分页位置标识。
	Marker         *string `json:"marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesResponse", string(data)}, " ")
}

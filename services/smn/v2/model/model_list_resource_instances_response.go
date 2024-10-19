package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceInstancesResponse Response Object
type ListResourceInstancesResponse struct {

	// 返回的资源列表。
	Resources *[]TagResource `json:"resources,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTagsResponse Response Object
type ListProjectTagsResponse struct {

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectTagsResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTagsResponse Response Object
type ListProjectTagsResponse struct {

	// 资源标签列表。
	Tags           *[]ResourceTags `json:"tags,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectTagsResponse", string(data)}, " ")
}

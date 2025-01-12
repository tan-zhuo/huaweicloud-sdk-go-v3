package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesByTagsResponse Response Object
type ListResourcesByTagsResponse struct {

	// 资源列表。
	Resources *[]ResourceDto `json:"resources,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListResourcesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsResponse", string(data)}, " ")
}

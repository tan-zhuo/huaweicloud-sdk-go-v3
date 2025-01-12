package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagsOfTenantResponse Response Object
type ListTagsOfTenantResponse struct {

	// 标签列表。
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTagsOfTenantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsOfTenantResponse struct{}"
	}

	return strings.Join([]string{"ListTagsOfTenantResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListListenerTagsResponse Response Object
type ListListenerTagsResponse struct {

	// 标签列表
	Tags           *[]ListTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"ListListenerTagsResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImagesTagsResponse Response Object
type ListImagesTagsResponse struct {

	// 标签列表
	Tags           *[]Tags `json:"tags,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListImagesTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImagesTagsResponse struct{}"
	}

	return strings.Join([]string{"ListImagesTagsResponse", string(data)}, " ")
}

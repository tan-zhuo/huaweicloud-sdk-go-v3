package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListResourcesByTagsRequestBody struct {

	// 包含标签。
	Tags []CreateTag `json:"tags"`
}

func (o ListResourcesByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsRequestBody", string(data)}, " ")
}

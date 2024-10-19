package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateResourceTagsRequestBody struct {

	// 标签列表。
	Tags *[]Tags `json:"tags,omitempty"`
}

func (o BatchCreateResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTagsRequestBody", string(data)}, " ")
}

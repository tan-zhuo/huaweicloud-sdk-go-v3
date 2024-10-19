package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreatePortTagsRequestBody This is a auto create Body Object
type BatchCreatePortTagsRequestBody struct {

	// 标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o BatchCreatePortTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePortTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreatePortTagsRequestBody", string(data)}, " ")
}

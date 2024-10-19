package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsVo struct {

	// 标签id
	TagId *string `json:"tag_id,omitempty"`

	// 标签键
	TagKey *string `json:"tag_key,omitempty"`

	// 标签值
	TagValue *string `json:"tag_value,omitempty"`
}

func (o TagsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsVo struct{}"
	}

	return strings.Join([]string{"TagsVo", string(data)}, " ")
}

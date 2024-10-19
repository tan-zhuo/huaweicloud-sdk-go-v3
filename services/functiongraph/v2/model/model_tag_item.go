package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagItem struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 值
	Values *[]string `json:"values,omitempty"`
}

func (o TagItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagItem struct{}"
	}

	return strings.Join([]string{"TagItem", string(data)}, " ")
}

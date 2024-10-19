package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindInstanceTags struct {

	// 实例标签
	Tags []Tag `json:"tags"`
}

func (o BindInstanceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindInstanceTags struct{}"
	}

	return strings.Join([]string{"BindInstanceTags", string(data)}, " ")
}

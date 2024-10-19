package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryImageByTagsResourceDetail 镜像详情
type QueryImageByTagsResourceDetail struct {

	// 镜像状态
	Status string `json:"status"`
}

func (o QueryImageByTagsResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryImageByTagsResourceDetail struct{}"
	}

	return strings.Join([]string{"QueryImageByTagsResourceDetail", string(data)}, " ")
}

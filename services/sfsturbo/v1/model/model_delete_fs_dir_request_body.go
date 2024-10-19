package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFsDirRequestBody 删除目录请求
type DeleteFsDirRequestBody struct {

	// 合法的目录全路径
	Path string `json:"path"`
}

func (o DeleteFsDirRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFsDirRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteFsDirRequestBody", string(data)}, " ")
}

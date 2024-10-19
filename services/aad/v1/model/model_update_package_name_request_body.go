package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePackageNameRequestBody 更新防护包名的请求体
type UpdatePackageNameRequestBody struct {

	// 名字
	Name string `json:"name"`
}

func (o UpdatePackageNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePackageNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePackageNameRequestBody", string(data)}, " ")
}

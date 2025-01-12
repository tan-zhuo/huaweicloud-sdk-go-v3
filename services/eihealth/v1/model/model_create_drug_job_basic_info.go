package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugJobBasicInfo 创建药物作业基本信息
type CreateDrugJobBasicInfo struct {

	// 作业的名称，取值范围：[5,64]，允许大小写字母、数字、空格、下划线(_)和中划线(-),只能以数字或字母开头
	Name string `json:"name"`

	// 标签，取值范围[0,5]，单个标签最大长度32字符，支持中文、字母、数字、空格、下划线和中划线，且不能以空格开头或者结尾。
	Labels *[]string `json:"labels,omitempty"`
}

func (o CreateDrugJobBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugJobBasicInfo struct{}"
	}

	return strings.Join([]string{"CreateDrugJobBasicInfo", string(data)}, " ")
}

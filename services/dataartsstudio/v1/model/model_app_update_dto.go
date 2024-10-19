package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppUpdateDto struct {

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`
}

func (o AppUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppUpdateDto struct{}"
	}

	return strings.Join([]string{"AppUpdateDto", string(data)}, " ")
}

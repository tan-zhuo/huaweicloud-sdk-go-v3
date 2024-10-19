package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TempName struct {

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o TempName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TempName struct{}"
	}

	return strings.Join([]string{"TempName", string(data)}, " ")
}

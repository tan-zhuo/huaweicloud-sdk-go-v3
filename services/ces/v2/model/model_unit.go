package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Unit 单位
type Unit struct {
}

func (o Unit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Unit struct{}"
	}

	return strings.Join([]string{"Unit", string(data)}, " ")
}

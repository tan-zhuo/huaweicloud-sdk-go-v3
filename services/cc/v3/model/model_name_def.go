package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NameDef 实例名字。
type NameDef struct {
}

func (o NameDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameDef struct{}"
	}

	return strings.Join([]string{"NameDef", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MysqlFlavorInfo flavor规格信息。
type MysqlFlavorInfo struct {

	// CPU核数。
	Vcpus string `json:"vcpus"`

	// 内存大小，单位GB。
	Ram string `json:"ram"`
}

func (o MysqlFlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlFlavorInfo struct{}"
	}

	return strings.Join([]string{"MysqlFlavorInfo", string(data)}, " ")
}

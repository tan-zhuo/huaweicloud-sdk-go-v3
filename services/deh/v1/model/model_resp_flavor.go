package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RespFlavor 弹性云服务器规格信息。
type RespFlavor struct {

	// 弹性云服务器规格ID。
	Id string `json:"id"`
}

func (o RespFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespFlavor struct{}"
	}

	return strings.Join([]string{"RespFlavor", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootBody 重启裸金属服务器接口请求结构体
type RebootBody struct {
	Reboot *ServersInfoType `json:"reboot"`
}

func (o RebootBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootBody struct{}"
	}

	return strings.Join([]string{"RebootBody", string(data)}, " ")
}

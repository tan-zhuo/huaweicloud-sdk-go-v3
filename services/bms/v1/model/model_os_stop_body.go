package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsStopBody 关闭裸金属服务器接口请求结构体
type OsStopBody struct {
	OsStop *OsStopBodyType `json:"os-stop"`
}

func (o OsStopBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsStopBody struct{}"
	}

	return strings.Join([]string{"OsStopBody", string(data)}, " ")
}

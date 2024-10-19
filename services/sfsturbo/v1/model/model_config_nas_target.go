package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigNasTarget 配置 nas 后端信息
type ConfigNasTarget struct {

	// nas 配置名
	Name string `json:"name"`

	// nas 配置协议类型
	Type string `json:"type"`

	// nas 配置 ip
	Url string `json:"url"`
}

func (o ConfigNasTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigNasTarget struct{}"
	}

	return strings.Join([]string{"ConfigNasTarget", string(data)}, " ")
}

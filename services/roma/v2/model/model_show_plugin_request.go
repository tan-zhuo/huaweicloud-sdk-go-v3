package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginRequest Request Object
type ShowPluginRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 插件编号
	PluginId string `json:"plugin_id"`
}

func (o ShowPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginRequest struct{}"
	}

	return strings.Join([]string{"ShowPluginRequest", string(data)}, " ")
}

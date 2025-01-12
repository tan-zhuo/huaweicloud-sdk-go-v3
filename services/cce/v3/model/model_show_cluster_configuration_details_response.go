package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterConfigurationDetailsResponse Response Object
type ShowClusterConfigurationDetailsResponse struct {

	// 获取指定集群配置项列表返回体
	Body           map[string][]PackageOptions `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowClusterConfigurationDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterConfigurationDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterConfigurationDetailsResponse", string(data)}, " ")
}

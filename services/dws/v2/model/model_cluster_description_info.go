package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterDescriptionInfo 集群描述对象
type ClusterDescriptionInfo struct {

	// 集群描述信息
	DescriptionInfo string `json:"description_info"`
}

func (o ClusterDescriptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDescriptionInfo struct{}"
	}

	return strings.Join([]string{"ClusterDescriptionInfo", string(data)}, " ")
}

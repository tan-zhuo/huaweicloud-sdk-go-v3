package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigMapsResponse Response Object
type ListConfigMapsResponse struct {

	// 配置项
	Configmaps *[]ConfigMapResp `json:"configmaps,omitempty"`

	// 满足条件的个数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListConfigMapsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigMapsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigMapsResponse", string(data)}, " ")
}

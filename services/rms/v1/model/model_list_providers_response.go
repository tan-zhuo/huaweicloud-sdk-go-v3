package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProvidersResponse Response Object
type ListProvidersResponse struct {

	// 云服务详情列表
	ResourceProviders *[]ResourceProviderResponse `json:"resource_providers,omitempty"`

	// 当前支持的云服务总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProvidersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProvidersResponse struct{}"
	}

	return strings.Join([]string{"ListProvidersResponse", string(data)}, " ")
}

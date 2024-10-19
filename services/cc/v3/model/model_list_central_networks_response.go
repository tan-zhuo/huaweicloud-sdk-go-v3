package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworksResponse Response Object
type ListCentralNetworksResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 中心网络列表。
	CentralNetworks []CentralNetwork `json:"central_networks"`
	HttpStatusCode  int              `json:"-"`
}

func (o ListCentralNetworksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworksResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworksResponse", string(data)}, " ")
}

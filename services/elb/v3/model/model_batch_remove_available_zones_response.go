package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveAvailableZonesResponse Response Object
type BatchRemoveAvailableZonesResponse struct {
	Loadbalancer *LoadBalancer `json:"loadbalancer,omitempty"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchRemoveAvailableZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveAvailableZonesResponse struct{}"
	}

	return strings.Join([]string{"BatchRemoveAvailableZonesResponse", string(data)}, " ")
}

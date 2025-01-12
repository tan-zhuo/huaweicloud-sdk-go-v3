package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaListAvailabilityZonesResponse Response Object
type NovaListAvailabilityZonesResponse struct {

	// 可用域信息。
	AvailabilityZoneInfo *[]NovaAvailabilityZone `json:"availabilityZoneInfo,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o NovaListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"NovaListAvailabilityZonesResponse", string(data)}, " ")
}

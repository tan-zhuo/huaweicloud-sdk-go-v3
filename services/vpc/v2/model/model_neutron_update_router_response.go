package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateRouterResponse Response Object
type NeutronUpdateRouterResponse struct {
	Router         *NeutronRouter `json:"router,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o NeutronUpdateRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateRouterResponse struct{}"
	}

	return strings.Join([]string{"NeutronUpdateRouterResponse", string(data)}, " ")
}

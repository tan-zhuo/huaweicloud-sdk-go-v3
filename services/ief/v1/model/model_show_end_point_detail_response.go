package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEndPointDetailResponse Response Object
type ShowEndPointDetailResponse struct {
	Endpoint       *EndpointObjResp `json:"endpoint,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowEndPointDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndPointDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowEndPointDetailResponse", string(data)}, " ")
}

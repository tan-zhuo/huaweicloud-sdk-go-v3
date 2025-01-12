package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServiceDetailResponse Response Object
type ShowServiceDetailResponse struct {
	Service        *ServiceRespDetail `json:"service,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowServiceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServiceDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowServiceDetailResponse", string(data)}, " ")
}

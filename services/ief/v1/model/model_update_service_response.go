package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServiceResponse Response Object
type UpdateServiceResponse struct {
	Service        *ServiceRespDetail `json:"service,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceResponse struct{}"
	}

	return strings.Join([]string{"UpdateServiceResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcRouteRequest Request Object
type CreateVpcRouteRequest struct {
	Body *CreateVpcRouteRequestBody `json:"body,omitempty"`
}

func (o CreateVpcRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcRouteRequest struct{}"
	}

	return strings.Join([]string{"CreateVpcRouteRequest", string(data)}, " ")
}

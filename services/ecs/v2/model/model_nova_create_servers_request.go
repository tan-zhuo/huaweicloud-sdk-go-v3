package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaCreateServersRequest Request Object
type NovaCreateServersRequest struct {

	// 微版本头
	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty"`

	Body *NovaCreateServersRequestBody `json:"body,omitempty"`
}

func (o NovaCreateServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateServersRequest struct{}"
	}

	return strings.Join([]string{"NovaCreateServersRequest", string(data)}, " ")
}

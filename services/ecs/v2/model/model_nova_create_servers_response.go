package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaCreateServersResponse Response Object
type NovaCreateServersResponse struct {
	Server         *NovaCreateServersResult `json:"server,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o NovaCreateServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateServersResponse struct{}"
	}

	return strings.Join([]string{"NovaCreateServersResponse", string(data)}, " ")
}

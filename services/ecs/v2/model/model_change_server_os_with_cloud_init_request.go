package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeServerOsWithCloudInitRequest Request Object
type ChangeServerOsWithCloudInitRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *ChangeServerOsWithCloudInitRequestBody `json:"body,omitempty"`
}

func (o ChangeServerOsWithCloudInitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithCloudInitRequest struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithCloudInitRequest", string(data)}, " ")
}

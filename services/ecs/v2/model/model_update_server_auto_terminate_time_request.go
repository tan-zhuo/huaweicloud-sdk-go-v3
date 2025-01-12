package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerAutoTerminateTimeRequest Request Object
type UpdateServerAutoTerminateTimeRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *UpdateServerAutoTerminateTimeRequestBody `json:"body,omitempty"`
}

func (o UpdateServerAutoTerminateTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerAutoTerminateTimeRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerAutoTerminateTimeRequest", string(data)}, " ")
}

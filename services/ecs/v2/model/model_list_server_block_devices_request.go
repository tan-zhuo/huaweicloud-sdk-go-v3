package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerBlockDevicesRequest Request Object
type ListServerBlockDevicesRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ListServerBlockDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerBlockDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListServerBlockDevicesRequest", string(data)}, " ")
}

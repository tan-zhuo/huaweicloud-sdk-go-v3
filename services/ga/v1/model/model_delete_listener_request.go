package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteListenerRequest Request Object
type DeleteListenerRequest struct {

	// 监听器ID。
	ListenerId string `json:"listener_id"`
}

func (o DeleteListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerRequest struct{}"
	}

	return strings.Join([]string{"DeleteListenerRequest", string(data)}, " ")
}

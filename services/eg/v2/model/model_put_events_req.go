package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutEventsReq struct {

	// putevent请求事件
	Events *[]CloudEvents `json:"events,omitempty"`
}

func (o PutEventsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutEventsReq struct{}"
	}

	return strings.Join([]string{"PutEventsReq", string(data)}, " ")
}

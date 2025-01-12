package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowQueueResponse Response Object
type BatchShowQueueResponse struct {

	// 队列信息列表。
	Queues *[]QueryQueueBase `json:"queues,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o BatchShowQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowQueueResponse struct{}"
	}

	return strings.Join([]string{"BatchShowQueueResponse", string(data)}, " ")
}

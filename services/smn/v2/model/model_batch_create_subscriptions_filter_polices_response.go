package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSubscriptionsFilterPolicesResponse Response Object
type BatchCreateSubscriptionsFilterPolicesResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 批量结果
	BatchResult    *[]BatchResult `json:"batch_result,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o BatchCreateSubscriptionsFilterPolicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubscriptionsFilterPolicesResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSubscriptionsFilterPolicesResponse", string(data)}, " ")
}

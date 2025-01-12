package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionsByTopicResponse Response Object
type ListSubscriptionsByTopicResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 订阅者个数。
	SubscriptionCount *int32 `json:"subscription_count,omitempty"`

	// Subscription结构体。
	Subscriptions  *[]ListSubscriptionsItem `json:"subscriptions,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListSubscriptionsByTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsByTopicResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsByTopicResponse", string(data)}, " ")
}

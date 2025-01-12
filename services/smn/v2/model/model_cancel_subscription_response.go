package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelSubscriptionResponse Response Object
type CancelSubscriptionResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"CancelSubscriptionResponse", string(data)}, " ")
}

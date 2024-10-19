package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceOrderResponse Response Object
type ChangeInstanceOrderResponse struct {

	// 订单ID。
	OrderId        *string `json:"orderId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeInstanceOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceOrderResponse struct{}"
	}

	return strings.Join([]string{"ChangeInstanceOrderResponse", string(data)}, " ")
}

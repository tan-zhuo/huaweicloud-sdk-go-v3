package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceChangeOrderReq struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 产品编号
	ProductId *string `json:"product_id,omitempty"`

	ResizeInfo *ResizeInstanceReq `json:"resize_info,omitempty"`
}

func (o InstanceChangeOrderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceChangeOrderReq struct{}"
	}

	return strings.Join([]string{"InstanceChangeOrderReq", string(data)}, " ")
}

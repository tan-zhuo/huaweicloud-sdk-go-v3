package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpcRequest Request Object
type ShowVpcRequest struct {

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`
}

func (o ShowVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcRequest struct{}"
	}

	return strings.Join([]string{"ShowVpcRequest", string(data)}, " ")
}

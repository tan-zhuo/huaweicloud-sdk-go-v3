package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcRequest Request Object
type DeleteVpcRequest struct {

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`
}

func (o DeleteVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcRequest", string(data)}, " ")
}

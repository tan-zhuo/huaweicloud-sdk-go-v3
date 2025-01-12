package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubnetRequest Request Object
type UpdateSubnetRequest struct {

	// 子网ID。
	SubnetId string `json:"subnet_id"`

	Body *UpdateSubnetRequestBody `json:"body,omitempty"`
}

func (o UpdateSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubnetRequest", string(data)}, " ")
}

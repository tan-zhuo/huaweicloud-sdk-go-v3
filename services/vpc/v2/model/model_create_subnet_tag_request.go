package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubnetTagRequest Request Object
type CreateSubnetTagRequest struct {

	// 子网ID
	SubnetId string `json:"subnet_id"`

	Body *CreateSubnetTagRequestBody `json:"body,omitempty"`
}

func (o CreateSubnetTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetTagRequest struct{}"
	}

	return strings.Join([]string{"CreateSubnetTagRequest", string(data)}, " ")
}

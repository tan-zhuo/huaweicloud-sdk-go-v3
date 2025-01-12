package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcRequest Request Object
type CreateVpcRequest struct {
	Body *CreateVpcRequestBody `json:"body,omitempty"`
}

func (o CreateVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcRequest struct{}"
	}

	return strings.Join([]string{"CreateVpcRequest", string(data)}, " ")
}

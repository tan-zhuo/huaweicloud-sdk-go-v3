package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubnetRequest Request Object
type CreateSubnetRequest struct {
	Body *CreateSubnetRequestBody `json:"body,omitempty"`
}

func (o CreateSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetRequest struct{}"
	}

	return strings.Join([]string{"CreateSubnetRequest", string(data)}, " ")
}

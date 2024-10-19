package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceRequest Request Object
type CreateInstanceRequest struct {
	Body *CreateInstanceReq `json:"body,omitempty"`
}

func (o CreateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequest", string(data)}, " ")
}

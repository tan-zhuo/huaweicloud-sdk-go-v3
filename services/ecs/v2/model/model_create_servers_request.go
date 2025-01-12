package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServersRequest Request Object
type CreateServersRequest struct {

	// 保证客户端请求幂等性的标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateServersRequestBody `json:"body,omitempty"`
}

func (o CreateServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServersRequest struct{}"
	}

	return strings.Join([]string{"CreateServersRequest", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeInstanceRequest Request Object
type ResizeInstanceRequest struct {
	Body *ChangeInstanceRequestBody `json:"body,omitempty"`
}

func (o ResizeInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceRequest struct{}"
	}

	return strings.Join([]string{"ResizeInstanceRequest", string(data)}, " ")
}

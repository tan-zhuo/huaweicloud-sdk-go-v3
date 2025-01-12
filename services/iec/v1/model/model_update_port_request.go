package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePortRequest Request Object
type UpdatePortRequest struct {

	// 端口ID
	PortId string `json:"port_id"`

	Body *UpdatePortRequestBody `json:"body,omitempty"`
}

func (o UpdatePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortRequest struct{}"
	}

	return strings.Join([]string{"UpdatePortRequest", string(data)}, " ")
}

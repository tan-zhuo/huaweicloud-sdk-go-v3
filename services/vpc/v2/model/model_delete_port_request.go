package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePortRequest Request Object
type DeletePortRequest struct {

	// 端口ID
	PortId string `json:"port_id"`
}

func (o DeletePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePortRequest struct{}"
	}

	return strings.Join([]string{"DeletePortRequest", string(data)}, " ")
}

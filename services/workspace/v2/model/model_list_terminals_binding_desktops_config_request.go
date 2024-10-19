package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTerminalsBindingDesktopsConfigRequest Request Object
type ListTerminalsBindingDesktopsConfigRequest struct {
}

func (o ListTerminalsBindingDesktopsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTerminalsBindingDesktopsConfigRequest struct{}"
	}

	return strings.Join([]string{"ListTerminalsBindingDesktopsConfigRequest", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSubNetworkInterfacesQuantityRequest Request Object
type ShowSubNetworkInterfacesQuantityRequest struct {
}

func (o ShowSubNetworkInterfacesQuantityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubNetworkInterfacesQuantityRequest struct{}"
	}

	return strings.Join([]string{"ShowSubNetworkInterfacesQuantityRequest", string(data)}, " ")
}

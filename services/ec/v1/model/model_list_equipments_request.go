package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEquipmentsRequest Request Object
type ListEquipmentsRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`
}

func (o ListEquipmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEquipmentsRequest struct{}"
	}

	return strings.Join([]string{"ListEquipmentsRequest", string(data)}, " ")
}

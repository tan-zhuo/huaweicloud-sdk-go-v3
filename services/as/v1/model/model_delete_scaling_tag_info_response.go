package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScalingTagInfoResponse Response Object
type DeleteScalingTagInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingTagInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingTagInfoResponse struct{}"
	}

	return strings.Join([]string{"DeleteScalingTagInfoResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLiveDataApiV2Request Request Object
type DeleteLiveDataApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 后端API的编号
	LdApiId string `json:"ld_api_id"`
}

func (o DeleteLiveDataApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLiveDataApiV2Request struct{}"
	}

	return strings.Join([]string{"DeleteLiveDataApiV2Request", string(data)}, " ")
}

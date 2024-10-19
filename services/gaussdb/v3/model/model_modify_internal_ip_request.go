package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyInternalIpRequest struct {

	// 内网IP。
	InternalIp string `json:"internal_ip"`
}

func (o ModifyInternalIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInternalIpRequest struct{}"
	}

	return strings.Join([]string{"ModifyInternalIpRequest", string(data)}, " ")
}

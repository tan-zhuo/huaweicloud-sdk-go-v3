package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindEipRequestBody struct {

	// 操作标识。取值： - BIND，表示绑定弹性公网IP。 - UNBIND，表示解绑弹性公网IP。
	Action string `json:"action"`

	// 弹性公网IP
	PublicIp string `json:"public_ip"`

	// 弹性公网IP的ID
	PublicIpId string `json:"public_ip_id"`
}

func (o BindEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindEipRequestBody struct{}"
	}

	return strings.Join([]string{"BindEipRequestBody", string(data)}, " ")
}

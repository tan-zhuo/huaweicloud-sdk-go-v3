package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerCidr 容器网络网段，指定cidrs字段使用时必填。
type ContainerCidr struct {

	// 容器网络网段，建议使用网段10.0.0.0/12~19，172.16.0.0/16~19，192.168.0.0/16~19。
	Cidr string `json:"cidr"`
}

func (o ContainerCidr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerCidr struct{}"
	}

	return strings.Join([]string{"ContainerCidr", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PostPaidServerSecurityGroup
type PostPaidServerSecurityGroup struct {

	// 可以为空，待创建云服务器的安全组，会对创建云服务器中配置的网卡生效。需要指定已有安全组的ID，UUID格式；若不传值，底层会按照空处理，不会创建安全组。
	Id *string `json:"id,omitempty"`
}

func (o PostPaidServerSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerSecurityGroup struct{}"
	}

	return strings.Join([]string{"PostPaidServerSecurityGroup", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainV2Request Request Object
type UpdateDomainV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	// 域名的编号
	DomainId string `json:"domain_id"`

	Body *UrlDomainModify `json:"body,omitempty"`
}

func (o UpdateDomainV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainV2Request struct{}"
	}

	return strings.Join([]string{"UpdateDomainV2Request", string(data)}, " ")
}

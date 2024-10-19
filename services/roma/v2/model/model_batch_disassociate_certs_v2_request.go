package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisassociateCertsV2Request Request Object
type BatchDisassociateCertsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	// 域名的编号
	DomainId string `json:"domain_id"`

	Body *AttachOrDetachCertsReqBody `json:"body,omitempty"`
}

func (o BatchDisassociateCertsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociateCertsV2Request struct{}"
	}

	return strings.Join([]string{"BatchDisassociateCertsV2Request", string(data)}, " ")
}

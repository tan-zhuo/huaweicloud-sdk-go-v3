package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePublicIpRequestBody 创建弹性公网IP请求体。
type CreatePublicIpRequestBody struct {
	Publicip *CreatePublicIpOption `json:"publicip"`

	Bandwidth *CreatePublicIpBandwidthOption `json:"bandwidth,omitempty"`
}

func (o CreatePublicIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicIpRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePublicIpRequestBody", string(data)}, " ")
}

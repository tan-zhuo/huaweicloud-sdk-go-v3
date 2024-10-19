package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDnsNameReq 申请内网域名
type CreateDnsNameReq struct {

	// 域名类型，当前只支持private。
	DnsType string `json:"dns_type"`
}

func (o CreateDnsNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDnsNameReq struct{}"
	}

	return strings.Join([]string{"CreateDnsNameReq", string(data)}, " ")
}

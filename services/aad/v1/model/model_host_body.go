package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostBody struct {

	// 域名
	DomainName string `json:"domain_name"`

	// 企业项目id，与接入的高防实例所属企业项目保持一致。可在华为云EPS服务中查看企业项目id，default企业项目id为\"0\"。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 高防实例ip列表。多个高防实例ip必须属于同一企业项目。
	Vips []string `json:"vips"`

	// 源站类型。 0 - 源站IP， 1 - 源站域名。
	RealServerType int32 `json:"real_server_type"`

	// HTTP端口，与port_https不能同时为空。DDoS高防支持的HTTP端口可在控制台查看。
	PortHttp *[]int32 `json:"port_http,omitempty"`

	// HTTPS端口，与port_http不能同时为空。DDoS高防支持的HTTPS端口可在控制台查看。
	PortHttps *[]int32 `json:"port_https,omitempty"`

	// 源站ip/源站域名
	RealServer string `json:"real_server"`
}

func (o HostBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostBody struct{}"
	}

	return strings.Join([]string{"HostBody", string(data)}, " ")
}

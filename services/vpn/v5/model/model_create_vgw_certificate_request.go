package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVgwCertificateRequest Request Object
type CreateVgwCertificateRequest struct {

	// VPN网关实例ID
	VgwId string `json:"vgw_id"`

	Body *CreateVpnGatewayCertificateRequestBody `json:"body,omitempty"`
}

func (o CreateVgwCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVgwCertificateRequest struct{}"
	}

	return strings.Join([]string{"CreateVgwCertificateRequest", string(data)}, " ")
}

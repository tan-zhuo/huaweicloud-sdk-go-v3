package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertificateRequest Request Object
type ExportCertificateRequest struct {

	// 证书id。
	CertificateId string `json:"certificate_id"`
}

func (o ExportCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateRequest struct{}"
	}

	return strings.Join([]string{"ExportCertificateRequest", string(data)}, " ")
}

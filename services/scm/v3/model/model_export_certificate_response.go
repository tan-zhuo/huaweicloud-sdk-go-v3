package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertificateResponse Response Object
type ExportCertificateResponse struct {

	// 证书及证书链。
	EntireCertificate *string `json:"entire_certificate,omitempty"`

	// 证书内容，不包含证书链。
	Certificate *string `json:"certificate,omitempty"`

	// 证书链。
	CertificateChain *string `json:"certificate_chain,omitempty"`

	// 证书私钥。
	PrivateKey *string `json:"private_key,omitempty"`

	// 国密证书返回，加密证书内容。
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// 国密证书返回，加密证书私钥。
	EncPrivateKey  *string `json:"enc_private_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateResponse struct{}"
	}

	return strings.Join([]string{"ExportCertificateResponse", string(data)}, " ")
}

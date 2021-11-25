package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ApplyCertificateToHostRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// https证书id，您可以通过调用查询证书列表（ListCertificates）接口获取证书id

	CertificateId string `json:"certificate_id"`

	Body *ApplyCertificateToHostRequestBody `json:"body,omitempty"`
}

func (o ApplyCertificateToHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyCertificateToHostRequest struct{}"
	}

	return strings.Join([]string{"ApplyCertificateToHostRequest", string(data)}, " ")
}

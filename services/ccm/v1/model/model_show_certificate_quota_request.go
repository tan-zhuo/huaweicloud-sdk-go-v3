package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateQuotaRequest Request Object
type ShowCertificateQuotaRequest struct {
}

func (o ShowCertificateQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificateQuotaRequest", string(data)}, " ")
}

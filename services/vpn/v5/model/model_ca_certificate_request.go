package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CaCertificateRequest 对端网关CA证书
type CaCertificateRequest struct {

	// 使用已有证书ID
	Id *string `json:"id,omitempty"`

	// 对端网关CA证书内容
	Content *string `json:"content,omitempty"`
}

func (o CaCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaCertificateRequest struct{}"
	}

	return strings.Join([]string{"CaCertificateRequest", string(data)}, " ")
}

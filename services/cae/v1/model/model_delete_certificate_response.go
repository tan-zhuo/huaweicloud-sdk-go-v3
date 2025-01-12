package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCertificateResponse Response Object
type DeleteCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeleteCertificateResponse", string(data)}, " ")
}

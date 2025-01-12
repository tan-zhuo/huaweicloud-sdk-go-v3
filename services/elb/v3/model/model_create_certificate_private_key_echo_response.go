package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificatePrivateKeyEchoResponse Response Object
type CreateCertificatePrivateKeyEchoResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	// 证书回显开关，项目粒度的,默认情况下,\"private_key_echo\"是true,证书的返回体中展示私钥。 当值为false时,证书的返回体中不展示私钥。
	PrivateKeyEcho *bool `json:"private_key_echo,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateCertificatePrivateKeyEchoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificatePrivateKeyEchoResponse struct{}"
	}

	return strings.Join([]string{"CreateCertificatePrivateKeyEchoResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCertificateRequestBody struct {

	// 证书名称，证书名称只能由数字、字母、中划线、下划线和英文句点组成，长度不能超过256位字符
	Name string `json:"name"`

	// 证书文件，仅支持PEM格式的证书和私钥文件，且文件中的换行符应以\\n替换，如请求示例所示
	Content string `json:"content"`

	// 证书私钥，仅支持PEM格式的证书和私钥文件，且文件中的换行符应以\\n替换，如请求示例所示
	Key string `json:"key"`
}

func (o CreateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequestBody", string(data)}, " ")
}

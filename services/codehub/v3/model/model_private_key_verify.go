package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivateKeyVerify struct {

	// 仓库UUID(由CreateRepository接口返回)
	RepositoryUuid string `json:"repository_uuid"`

	// 私钥
	PrivateKey string `json:"private_key"`
}

func (o PrivateKeyVerify) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateKeyVerify struct{}"
	}

	return strings.Join([]string{"PrivateKeyVerify", string(data)}, " ")
}

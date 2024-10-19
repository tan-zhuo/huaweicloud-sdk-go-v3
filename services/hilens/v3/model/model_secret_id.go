package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecretId struct {

	// 密钥ID
	Id string `json:"id"`
}

func (o SecretId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecretId struct{}"
	}

	return strings.Join([]string{"SecretId", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSecretTagRequestBody struct {
	Tag *TagItem `json:"tag"`
}

func (o CreateSecretTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecretTagRequestBody", string(data)}, " ")
}

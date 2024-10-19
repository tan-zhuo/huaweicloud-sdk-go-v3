package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangePasswordBody struct {
	Password string `json:"password"`
}

func (o ChangePasswordBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordBody struct{}"
	}

	return strings.Join([]string{"ChangePasswordBody", string(data)}, " ")
}

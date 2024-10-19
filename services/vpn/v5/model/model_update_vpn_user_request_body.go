package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnUserRequestBody struct {
	User *UpdateVpnUserRequestBodyContent `json:"user"`
}

func (o UpdateVpnUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserRequestBody", string(data)}, " ")
}

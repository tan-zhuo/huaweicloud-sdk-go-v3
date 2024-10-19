package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserRequestBody
type UpdateUserRequestBody struct {
	User *UpdateUserOption `json:"user"`
}

func (o UpdateUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateUserRequestBody", string(data)}, " ")
}

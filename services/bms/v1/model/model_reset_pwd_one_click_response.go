package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetPwdOneClickResponse Response Object
type ResetPwdOneClickResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetPwdOneClickResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPwdOneClickResponse struct{}"
	}

	return strings.Join([]string{"ResetPwdOneClickResponse", string(data)}, " ")
}

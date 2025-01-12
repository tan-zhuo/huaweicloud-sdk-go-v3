package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetServerPasswordResponse Response Object
type ResetServerPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetServerPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetServerPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetServerPasswordResponse", string(data)}, " ")
}

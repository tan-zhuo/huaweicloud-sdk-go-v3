package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopNamePolicyResponse Response Object
type BatchDeleteDesktopNamePolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteDesktopNamePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopNamePolicyResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopNamePolicyResponse", string(data)}, " ")
}

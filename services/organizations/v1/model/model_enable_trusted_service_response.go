package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableTrustedServiceResponse Response Object
type EnableTrustedServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableTrustedServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableTrustedServiceResponse struct{}"
	}

	return strings.Join([]string{"EnableTrustedServiceResponse", string(data)}, " ")
}

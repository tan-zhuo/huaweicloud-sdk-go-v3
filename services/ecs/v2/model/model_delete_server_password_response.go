package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerPasswordResponse Response Object
type DeleteServerPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerPasswordResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerPasswordResponse", string(data)}, " ")
}

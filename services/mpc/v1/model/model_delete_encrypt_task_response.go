package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEncryptTaskResponse Response Object
type DeleteEncryptTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEncryptTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEncryptTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteEncryptTaskResponse", string(data)}, " ")
}

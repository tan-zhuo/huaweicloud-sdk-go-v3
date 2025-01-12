package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePermanentAccessKeyResponse Response Object
type DeletePermanentAccessKeyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePermanentAccessKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePermanentAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"DeletePermanentAccessKeyResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTokenResponse Response Object
type DeleteTokenResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTokenResponse struct{}"
	}

	return strings.Join([]string{"DeleteTokenResponse", string(data)}, " ")
}

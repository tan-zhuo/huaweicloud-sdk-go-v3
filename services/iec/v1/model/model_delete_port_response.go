package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePortResponse Response Object
type DeletePortResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePortResponse struct{}"
	}

	return strings.Join([]string{"DeletePortResponse", string(data)}, " ")
}

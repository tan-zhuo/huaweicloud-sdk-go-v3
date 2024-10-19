package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomIngressPortResponse Response Object
type DeleteCustomIngressPortResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomIngressPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomIngressPortResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomIngressPortResponse", string(data)}, " ")
}

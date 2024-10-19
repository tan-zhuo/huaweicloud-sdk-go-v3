package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainResponse Response Object
type DeleteDomainResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainResponse", string(data)}, " ")
}

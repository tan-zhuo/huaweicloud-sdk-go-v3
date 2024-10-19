package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainsResponse Response Object
type DeleteDomainsResponse struct {
	Data           *DomainSetResponseData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o DeleteDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainsResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainsResponse", string(data)}, " ")
}

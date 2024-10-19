package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListAuthDomainsRequest Request Object
type KeystoneListAuthDomainsRequest struct {
}

func (o KeystoneListAuthDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListAuthDomainsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListAuthDomainsRequest", string(data)}, " ")
}

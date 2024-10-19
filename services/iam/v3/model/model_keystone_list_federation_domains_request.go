package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListFederationDomainsRequest Request Object
type KeystoneListFederationDomainsRequest struct {
}

func (o KeystoneListFederationDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListFederationDomainsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListFederationDomainsRequest", string(data)}, " ")
}

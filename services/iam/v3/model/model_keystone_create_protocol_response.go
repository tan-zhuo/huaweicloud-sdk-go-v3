package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateProtocolResponse Response Object
type KeystoneCreateProtocolResponse struct {
	Protocol       *ProtocolResult `json:"protocol,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o KeystoneCreateProtocolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateProtocolResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateProtocolResponse", string(data)}, " ")
}

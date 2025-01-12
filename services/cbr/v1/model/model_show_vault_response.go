package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVaultResponse Response Object
type ShowVaultResponse struct {
	Vault          *Vault `json:"vault,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultResponse struct{}"
	}

	return strings.Join([]string{"ShowVaultResponse", string(data)}, " ")
}

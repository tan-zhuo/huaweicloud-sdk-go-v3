package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOpenIdConnectConfigResponse Response Object
type UpdateOpenIdConnectConfigResponse struct {
	OpenidConnectConfig *OpenIdConnectConfig `json:"openid_connect_config,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o UpdateOpenIdConnectConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOpenIdConnectConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateOpenIdConnectConfigResponse", string(data)}, " ")
}

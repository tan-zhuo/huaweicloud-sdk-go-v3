package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOpenIdConnectConfigRequest Request Object
type CreateOpenIdConnectConfigRequest struct {

	// 身份提供商ID
	IdpId string `json:"idp_id"`

	Body *CreateOpenIdConnectConfigRequestBody `json:"body,omitempty"`
}

func (o CreateOpenIdConnectConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOpenIdConnectConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateOpenIdConnectConfigRequest", string(data)}, " ")
}

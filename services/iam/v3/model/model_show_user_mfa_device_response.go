package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserMfaDeviceResponse Response Object
type ShowUserMfaDeviceResponse struct {
	VirtualMfaDevice *MfaDeviceResult `json:"virtual_mfa_device,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o ShowUserMfaDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserMfaDeviceResponse struct{}"
	}

	return strings.Join([]string{"ShowUserMfaDeviceResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMfaDeviceReq
type CreateMfaDeviceReq struct {
	VirtualMfaDevice *CreateMfaDevice `json:"virtual_mfa_device"`
}

func (o CreateMfaDeviceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMfaDeviceReq struct{}"
	}

	return strings.Join([]string{"CreateMfaDeviceReq", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateMappingRequestBody
type KeystoneUpdateMappingRequestBody struct {
	Mapping *MappingOption `json:"mapping"`
}

func (o KeystoneUpdateMappingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateMappingRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateMappingRequestBody", string(data)}, " ")
}

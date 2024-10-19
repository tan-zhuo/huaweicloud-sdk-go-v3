package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListMappingsRequest Request Object
type KeystoneListMappingsRequest struct {
}

func (o KeystoneListMappingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListMappingsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListMappingsRequest", string(data)}, " ")
}

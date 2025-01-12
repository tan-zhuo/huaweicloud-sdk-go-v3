package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CacheConfigRequestBody This is a auto create Body Object
type CacheConfigRequestBody struct {
	CacheConfig *CacheConfigRequest `json:"cache_config"`
}

func (o CacheConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CacheConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CacheConfigRequestBody", string(data)}, " ")
}

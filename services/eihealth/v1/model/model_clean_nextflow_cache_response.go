package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CleanNextflowCacheResponse Response Object
type CleanNextflowCacheResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CleanNextflowCacheResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CleanNextflowCacheResponse struct{}"
	}

	return strings.Join([]string{"CleanNextflowCacheResponse", string(data)}, " ")
}

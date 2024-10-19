package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePublicationsResponse Response Object
type BatchDeletePublicationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeletePublicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicationsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicationsResponse", string(data)}, " ")
}

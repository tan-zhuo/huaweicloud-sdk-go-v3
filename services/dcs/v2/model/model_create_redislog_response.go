package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRedislogResponse Response Object
type CreateRedislogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateRedislogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedislogResponse struct{}"
	}

	return strings.Join([]string{"CreateRedislogResponse", string(data)}, " ")
}

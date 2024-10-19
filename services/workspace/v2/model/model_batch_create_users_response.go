package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateUsersResponse Response Object
type BatchCreateUsersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateUsersResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateUsersResponse", string(data)}, " ")
}

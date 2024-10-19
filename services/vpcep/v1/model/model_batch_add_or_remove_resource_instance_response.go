package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddOrRemoveResourceInstanceResponse Response Object
type BatchAddOrRemoveResourceInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddOrRemoveResourceInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrRemoveResourceInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchAddOrRemoveResourceInstanceResponse", string(data)}, " ")
}

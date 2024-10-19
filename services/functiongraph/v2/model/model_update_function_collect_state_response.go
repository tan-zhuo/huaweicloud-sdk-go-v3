package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFunctionCollectStateResponse Response Object
type UpdateFunctionCollectStateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateFunctionCollectStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionCollectStateResponse struct{}"
	}

	return strings.Join([]string{"UpdateFunctionCollectStateResponse", string(data)}, " ")
}

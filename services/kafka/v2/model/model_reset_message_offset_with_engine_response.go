package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetMessageOffsetWithEngineResponse Response Object
type ResetMessageOffsetWithEngineResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetMessageOffsetWithEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetMessageOffsetWithEngineResponse struct{}"
	}

	return strings.Join([]string{"ResetMessageOffsetWithEngineResponse", string(data)}, " ")
}

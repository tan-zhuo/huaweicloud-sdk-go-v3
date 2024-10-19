package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnTagResourceResponse Response Object
type UnTagResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UnTagResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnTagResourceResponse struct{}"
	}

	return strings.Join([]string{"UnTagResourceResponse", string(data)}, " ")
}

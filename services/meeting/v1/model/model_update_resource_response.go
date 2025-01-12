package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceResponse Response Object
type UpdateResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceResponse", string(data)}, " ")
}

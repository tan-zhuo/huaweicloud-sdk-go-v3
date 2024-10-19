package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppResponse Response Object
type UpdateAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppResponse", string(data)}, " ")
}

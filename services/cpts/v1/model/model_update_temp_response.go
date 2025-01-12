package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTempResponse Response Object
type UpdateTempResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTempResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTempResponse struct{}"
	}

	return strings.Join([]string{"UpdateTempResponse", string(data)}, " ")
}

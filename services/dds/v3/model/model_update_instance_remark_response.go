package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRemarkResponse Response Object
type UpdateInstanceRemarkResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceRemarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRemarkResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRemarkResponse", string(data)}, " ")
}

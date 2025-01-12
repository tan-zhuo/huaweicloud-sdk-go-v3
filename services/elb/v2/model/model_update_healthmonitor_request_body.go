package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHealthmonitorRequestBody This is a auto create Body Object
type UpdateHealthmonitorRequestBody struct {
	Healthmonitor *UpdateHealthmonitorReq `json:"healthmonitor"`
}

func (o UpdateHealthmonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthmonitorRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHealthmonitorRequestBody", string(data)}, " ")
}

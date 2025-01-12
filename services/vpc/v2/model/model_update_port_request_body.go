package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePortRequestBody
type UpdatePortRequestBody struct {
	Port *UpdatePortOption `json:"port"`
}

func (o UpdatePortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePortRequestBody", string(data)}, " ")
}

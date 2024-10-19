package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EmptyJsonResponse struct {
}

func (o EmptyJsonResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmptyJsonResponse struct{}"
	}

	return strings.Join([]string{"EmptyJsonResponse", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Handler 备注信息，已处理的告警才有
type Handler struct {
}

func (o Handler) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Handler struct{}"
	}

	return strings.Join([]string{"Handler", string(data)}, " ")
}

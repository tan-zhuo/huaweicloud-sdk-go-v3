package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleTime 处理时间，毫秒，已处理的告警才有
type HandleTime struct {
}

func (o HandleTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleTime struct{}"
	}

	return strings.Join([]string{"HandleTime", string(data)}, " ")
}

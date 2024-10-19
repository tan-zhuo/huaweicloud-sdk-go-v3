package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndTime 屏蔽截止时间，HH:mm:ss。
type EndTime struct {
}

func (o EndTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndTime struct{}"
	}

	return strings.Join([]string{"EndTime", string(data)}, " ")
}

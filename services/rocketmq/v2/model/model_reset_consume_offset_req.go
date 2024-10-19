package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetConsumeOffsetReq struct {

	// 重置的主题。
	Topic string `json:"topic"`

	// 重置的时间。
	Timestamp string `json:"timestamp"`
}

func (o ResetConsumeOffsetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetReq struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetReq", string(data)}, " ")
}

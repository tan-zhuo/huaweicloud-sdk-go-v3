package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRes
type CreateRes struct {

	// 视频id
	Id string `json:"id"`
}

func (o CreateRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRes struct{}"
	}

	return strings.Join([]string{"CreateRes", string(data)}, " ")
}

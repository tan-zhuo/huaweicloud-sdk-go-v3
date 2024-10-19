package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupId 服务器组ID
type GroupId struct {
}

func (o GroupId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupId struct{}"
	}

	return strings.Join([]string{"GroupId", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 重启方式
type RebootType struct {

	// 重启方式
	Type string `json:"type"`
}

func (o RebootType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootType struct{}"
	}

	return strings.Join([]string{"RebootType", string(data)}, " ")
}
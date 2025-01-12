package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessPid 进程id
type ProcessPid struct {
}

func (o ProcessPid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessPid struct{}"
	}

	return strings.Join([]string{"ProcessPid", string(data)}, " ")
}

package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInstanceSpecRsp struct {

	// 实例详情响应体
	Containers *[]TaskInstanceSpecContainersRsp `json:"containers,omitempty"`
}

func (o TaskInstanceSpecRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInstanceSpecRsp struct{}"
	}

	return strings.Join([]string{"TaskInstanceSpecRsp", string(data)}, " ")
}

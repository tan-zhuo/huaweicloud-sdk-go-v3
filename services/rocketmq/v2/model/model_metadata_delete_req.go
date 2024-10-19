package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetadataDeleteReq struct {

	// 需要删除的任务列表。
	TaskIds []string `json:"task_ids"`
}

func (o MetadataDeleteReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetadataDeleteReq struct{}"
	}

	return strings.Join([]string{"MetadataDeleteReq", string(data)}, " ")
}

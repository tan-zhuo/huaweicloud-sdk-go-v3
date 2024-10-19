package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateMqsInstanceTopicReq struct {

	// Topic列表。
	Topics []UpdateTopicObject `json:"topics"`
}

func (o UpdateMqsInstanceTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMqsInstanceTopicReq struct{}"
	}

	return strings.Join([]string{"UpdateMqsInstanceTopicReq", string(data)}, " ")
}

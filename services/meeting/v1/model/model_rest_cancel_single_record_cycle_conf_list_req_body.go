package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestCancelSingleRecordCycleConfListReqBody struct {

	// 待删除的子会议UUID列表。
	CycleSubConfIDs []string `json:"cycleSubConfIDs"`
}

func (o RestCancelSingleRecordCycleConfListReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestCancelSingleRecordCycleConfListReqBody struct{}"
	}

	return strings.Join([]string{"RestCancelSingleRecordCycleConfListReqBody", string(data)}, " ")
}

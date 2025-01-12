package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlowLogReqBody
type UpdateFlowLogReqBody struct {
	FlowLog *UpdateFlowLogReq `json:"flow_log"`
}

func (o UpdateFlowLogReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlowLogReqBody struct{}"
	}

	return strings.Join([]string{"UpdateFlowLogReqBody", string(data)}, " ")
}
